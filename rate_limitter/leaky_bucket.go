package ratelimitter

import (
	"fmt"
	"time"
)

type Request struct {
	ID string
}

type LeakyBucket struct {
	queue    chan Request
	interval time.Duration
	stopChan chan struct{}
}

func NewLeakyBucket(capacity int, interval time.Duration) *LeakyBucket {
	lb := &LeakyBucket{
		queue:    make(chan Request, capacity),
		interval: interval,
		stopChan: make(chan struct{}),
	}

	go lb.StartLeaking()

	return lb
}

func (lb *LeakyBucket) TryAdd(req Request) {
	select {
	case lb.queue <- req:
		fmt.Printf("Request added %s\n", req.ID)
	default:
		fmt.Printf("Request not added %s\n", req.ID)
	}
}

func (lb *LeakyBucket) StartLeaking() {
	ticker := time.NewTicker(lb.interval)
	defer ticker.Stop()

	for {
		select {
		case <-lb.stopChan:
			return
		case <-ticker.C:
			select {
			case req := <-lb.queue:
				fmt.Printf("Processing request %s\n", req.ID)
			default:
			}
		}
	}
}

func main() {
	bucket := NewLeakyBucket(5, 1*time.Second) // max 5 requests in queue, 1 processed/sec

	for i := 1; i <= 10; i++ {
		bucket.TryAdd(Request{ID: i})
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(7 * time.Second)
	bucket.Stop()
}
