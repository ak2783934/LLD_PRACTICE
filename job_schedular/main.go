package main

import (
	"errors"
	"sync"
	"time"
)

type Job struct {
	JobID          string
	Payload        string
	cronExpression string
}

type JobExecutionStatus string

const (
	Enqueued   JobExecutionStatus = "ENQUEUED"
	Processing JobExecutionStatus = "PROCESSING"
	Retrying   JobExecutionStatus = "RETRYING"
	Completed  JobExecutionStatus = "COMPLETED"
	Cancelled  JobExecutionStatus = "CANCELLED"
	Failed     JobExecutionStatus = "FAILED"
)

type JobExecution struct {
	runID         string
	jobID         string
	executionTime time.Time
	retryCount    int
	status        JobExecutionStatus
}

// keeping the structure like this to make sure things are accessible faster
type JobSchedular struct {
	JobsRepository      map[string]*Job
	ExecutionRepository map[string][]*JobExecution
	MaxRetryCount       int
}

var once sync.Once
var JobSchedularInstance *JobSchedular

func CreateJobSchedular() *JobSchedular {
	once.Do(func() {
		JobSchedularInstance = &JobSchedular{
			JobsRepository:      make(map[string]*Job),
			ExecutionRepository: make(map[string][]*JobExecution),
			MaxRetryCount:       0,
		}
	})
	return JobSchedularInstance
}

func (J *JobSchedular) CreateJob(payload string, cronExp string) {
	job := &Job{
		JobID:          generateUUID(),
		Payload:        payload,
		cronExpression: cronExp,
	}
	J.JobsRepository[job.JobID] = job
	nextRunTime := getNextRunTime(cronExp)
	jobExecution := &JobExecution{
		runID:         generateUUID(),
		jobID:         job.JobID,
		executionTime: nextRunTime,
		retryCount:    0,
		status:        Enqueued,
	}
	J.ExecutionRepository[job.JobID] = append(J.ExecutionRepository[job.JobID], jobExecution)
}

func (J *JobSchedular) CancelJob(jobID string) error {
	_, ok := J.JobsRepository[jobID]
	if !ok {
		return errors.New("Job id invalid")
	}
	jobExecutions := J.ExecutionRepository[jobID]
	// assuming only last job must be pending or under execution, or else all other jobs are done.
	n := len(jobExecutions)
	lastJob := jobExecutions[n-1]

	if lastJob.status == Processing {
		return errors.New("Job already under processing")
	}

	if lastJob.status == Completed {
		return errors.New("Job already executed")
	}

	if lastJob.status == Cancelled {
		return errors.New("Job already cancelled")
	}

	if lastJob.status == Failed {
		return errors.New("Job already failed")
	}

	lastJob.status = Cancelled

	return nil
}


func (J *JobSchedular) RunCron(executor JobExecutor) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		now := time.Now()

		// iterate over all job executions
		for jobID, executions := range J.ExecutionRepository {
			for _, exec := range executions {

				// pick only due & enqueued executions
				if exec.status != Enqueued {
					continue
				}
				if exec.executionTime.After(now) {
					continue
				}

				// 1️⃣ Claim the execution
				exec.status = Processing

				// 2️⃣ Execute the job
				err := executor.Execute(exec)

				if err == nil {
					// 3️⃣ Success path
					exec.status = Completed

					// 4️⃣ Schedule next run if recurring
					job := J.JobsRepository[jobID]
					if job.cronExpression != "" {
						nextRun := getNextRunTime(job.cronExpression)

						nextExec := &JobExecution{
							runID:         generateUUID(),
							jobID:         jobID,
							executionTime: nextRun,
							retryCount:    0,
							status:        Enqueued,
						}
						J.ExecutionRepository[jobID] =
							append(J.ExecutionRepository[jobID], nextExec)
					}

				} else {
					// 5️⃣ Failure path
					exec.retryCount++

					if exec.retryCount <= J.MaxRetryCount {
						// retry with simple backoff
						exec.status = Retrying
						exec.executionTime = now.Add(time.Second * time.Duration(exec.retryCount))
						exec.status = Enqueued
					} else {
						exec.status = Failed
					}
				}
			}
		}
	}
}



type JobExecutor interface {
	Execute(*JobExecution) error
}

func generateUUID() string {
	return time.Now().Format("123123.123123")
}

func getNextRunTime(cronExp string) time.Time {
	return time.Time{}
}
