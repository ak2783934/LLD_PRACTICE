package main

import (
	"fmt"
	"sync"
	"time"
)

type Transaction struct {
	MerchantID string
	Amount     int64
	Timestamp  time.Time
}

type Aggregator struct {
	mu      sync.RWMutex
	buckets map[int64]map[string]int64 // minute -> merchant -> total
	retain  int64                      // number of minutes to retain
}

func NewAggregator(retainMinutes int64) *Aggregator {
	return &Aggregator{
		buckets: make(map[int64]map[string]int64),
		retain:  retainMinutes,
	}
}

func minuteBucket(t time.Time) int64 {
	return t.Unix() / 60
}

// Add ingests one transaction from the active stream.
func (a *Aggregator) Add(tx Transaction) {
	b := minuteBucket(tx.Timestamp)

	a.mu.Lock()
	defer a.mu.Unlock()

	if _, ok := a.buckets[b]; !ok {
		a.buckets[b] = make(map[string]int64)
	}
	a.buckets[b][tx.MerchantID] += tx.Amount

	a.evictLocked(time.Now())
}

// Query returns total amount per merchant for the last n minutes.
func (a *Aggregator) Query(n int64, now time.Time) map[string]int64 {
	result := make(map[string]int64)
	currentBucket := minuteBucket(now)
	startBucket := currentBucket - n + 1

	a.mu.RLock()
	defer a.mu.RUnlock()

	for b := startBucket; b <= currentBucket; b++ {
		merchantTotals, ok := a.buckets[b]
		if !ok {
			continue
		}
		for merchantID, amt := range merchantTotals {
			result[merchantID] += amt
		}
	}

	return result
}

func (a *Aggregator) evictLocked(now time.Time) {
	minAllowed := minuteBucket(now) - a.retain + 1
	for b := range a.buckets {
		if b < minAllowed {
			delete(a.buckets, b)
		}
	}
}

func main() {
	agg := NewAggregator(120) // keep last 120 minutes

	now := time.Now()

	agg.Add(Transaction{MerchantID: "m1", Amount: 100, Timestamp: now.Add(-2 * time.Minute)})
	agg.Add(Transaction{MerchantID: "m1", Amount: 50, Timestamp: now.Add(-1 * time.Minute)})
	agg.Add(Transaction{MerchantID: "m2", Amount: 200, Timestamp: now.Add(-1 * time.Minute)})
	agg.Add(Transaction{MerchantID: "m1", Amount: 30, Timestamp: now})

	fmt.Println("Last 1 min:", agg.Query(1, now))
	fmt.Println("Last 3 min:", agg.Query(3, now))
}
