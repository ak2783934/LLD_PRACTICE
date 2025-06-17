package ratelimitter

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	Capacity     int
	Token        int
	RefillRate   int
	LastFillTime time.Time
	Interval     time.Duration
	mutex        sync.Mutex
}

func NewTokenBucket(capacity int, refillRate int, interval time.Duration) *TokenBucket {
	return &TokenBucket{
		Capacity:     capacity,
		Token:        capacity,
		RefillRate:   refillRate,
		LastFillTime: time.Now(),
		Interval:     interval,
	}
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.mutex.Lock()
	defer tb.mutex.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(tb.LastFillTime)

	// then only we refill the bucket.
	if elapsedTime >= tb.Interval {
		amountToBeFilled := int(elapsedTime/tb.Interval) * tb.RefillRate
		if amountToBeFilled > 0 {
			tb.Token = min(tb.Capacity, tb.Token+amountToBeFilled)
			tb.LastFillTime = now
		}
	}

	if tb.Token > 0 {
		tb.Token--
		return true
	} else {
		return false
	}
}

func main() {
	bucket := NewTokenBucket(5, 1, 2*time.Second)

	for i := 0; i < 10; i++ {
		if bucket.AllowRequest() {
			fmt.Println("Request", i, "allowed at", time.Now())
		} else {
			fmt.Println("Request", i, "blocked at", time.Now())
		}
		time.Sleep(1 * time.Second)
	}
}
