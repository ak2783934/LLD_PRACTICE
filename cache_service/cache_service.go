package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Value struct {
	value     string
	ttlSecond int
	expire    time.Time
}

type CacheService struct {
	KeyValueStore map[string]*Value
	mu            sync.RWMutex
}

func NewCacheService() *CacheService {
	return &CacheService{
		KeyValueStore: make(map[string]*Value),
	}
}

func (c *CacheService) SET(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.KeyValueStore[key] = &Value{
		value:     value,
		ttlSecond: -1,
	}
}
func (c *CacheService) GET(key string) (string, error) {
	c.mu.RLock()
	value, ok := c.KeyValueStore[key]
	c.mu.RUnlock()
	if !ok {
		return "", errors.New("key not found")
	}
	if value.ttlSecond != -1 {
		if value.expire.Before(time.Now()) {
			c.mu.Lock()
			delete(c.KeyValueStore, key)
			c.mu.Unlock()
			return "", errors.New("key not found")
		}
	}

	return value.value, nil
}
func (c *CacheService) DEL(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.KeyValueStore, key)
}

// this is the only place where we are setting TTL to a key
func (c *CacheService) EXPIRE(key string, ttlSecond int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.KeyValueStore[key]
	if !ok {
		return errors.New("key doesn't exist")
	}

	// ttl already existed.
	if value.ttlSecond != -1 {
		if value.expire.Before(time.Now()) {
			delete(c.KeyValueStore, key)
			return errors.New("key doesn't exist")
		}
	}
	value.ttlSecond = ttlSecond
	value.expire = time.Now().Add(time.Second * time.Duration(ttlSecond))

	return nil
}

func (c *CacheService) TTL(key string) (int, error) {
	c.mu.RLock()
	value, ok := c.KeyValueStore[key]
	if !ok {
		return 0, errors.New("key doesn't exist")
	}
	c.mu.RUnlock()
	if value.expire.Before(time.Now()) {
		fmt.Println("key expired")
		c.mu.Lock()
		delete(c.KeyValueStore, key)
		c.mu.Unlock()
		return 0, errors.New("key doesn't exist")
	}

	return int(value.expire.Sub(time.Now()).Seconds()), nil
}

func main() {
	cacheService := NewCacheService()

	cacheService.SET("name", "Avinash")
	value, _ := cacheService.GET("name") // returns "Avinash"
	fmt.Println("value ", value)

	cacheService.EXPIRE("name", 5) // key expires in 5 seconds
	time.Sleep(6 * time.Second)
	ttl, _ := cacheService.TTL("name") // returns remaining TTL in seconds
	fmt.Println("TTL left", ttl)

	time.Sleep(5 * time.Second)
	value1, err := cacheService.GET("name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value1)

	// cacheService.DEL("name")
}
