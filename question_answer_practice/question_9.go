package main

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

/*
we have an array of payment jobs
we loop through the array and push them into channels.

we have a worker pool at the end of channel, those consume and do the validation.

once validation is done, they again push the jobs into the result queue.
*/

type PaymentValidationJob struct {
	ID         string
	PaymentID  string
	MerchantID string
	Amount     int
	Timestamp  time.Time
	Attempt    int
	MaxRetries int
}

type Result struct {
	JobID string
	Err   error
}

type WorkerPool struct {
	jobs    chan PaymentValidationJob
	results chan Result
	dlq     chan PaymentValidationJob
	wg      sync.WaitGroup
}

func NewWorkerPool(queueSize int) *WorkerPool {
	return &WorkerPool{
		jobs:    make(chan PaymentValidationJob, queueSize),
		results: make(chan Result, queueSize),
		dlq:     make(chan PaymentValidationJob, queueSize),
	}
}

func (wp *WorkerPool) AddJob(job PaymentValidationJob) {
	wp.jobs <- job
}

func (wp *WorkerPool) Start(ctx context.Context, worker int) {
	for i := range worker {
		wp.wg.Add(1)
		go wp.worker(ctx, i)
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
	close(wp.dlq)
}

func (wp *WorkerPool) worker(ctx context.Context, workerID int) {
	// here we consume the msg and do valiation and exponential backoff

	defer wp.wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-wp.jobs:
			if !ok {
				return
			}

			err := paymentValidation(job)
			if err == nil {
				wp.results <- Result{JobID: job.ID, Err: nil}
				continue
			}

			if isTransient(err) && job.Attempt < job.MaxRetries {
				job.Attempt++
				delay := backoffWithJitter(job.Attempt, 200*time.Millisecond, 5*time.Second)

				go func(job PaymentValidationJob, d time.Duration) {
					ticker := time.NewTimer(d)
					defer ticker.Stop()

					select {
					case <-ctx.Done():
						return
					case <-ticker.C:
						wp.AddJob(job)
					}
				}(job, delay)
			}

			if job.Attempt >= job.MaxRetries {
				wp.dlq <- job
			}
			wp.results <- Result{JobID: job.ID, Err: err}
		}
	}
}

func isTransient(err error) bool {
	return err != nil && err.Error() == "transient upstream validation failure"
}

func backoffWithJitter(attempt int, base, max time.Duration) time.Duration {
	delay := base * time.Duration(1<<(attempt-1))
	if delay > max {
		delay = max
	}

	jitter := time.Duration(rand.Int63n(int64(delay / 5))) // up to 20%
	return delay + jitter
}
