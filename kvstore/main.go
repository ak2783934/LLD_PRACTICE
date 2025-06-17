package main

import (
	"fmt"
	"sync"
)

type KVStore struct {
	mutex sync.RWMutex
	store map[string]string
}

func NewKVStore() *KVStore {
	return &KVStore{
		store: map[string]string{},
	}
}

func (k *KVStore) Get(key string) (string, bool) {
	k.mutex.RLock()
	defer k.mutex.Unlock()
	value, ok := k.store[key]
	return value, ok
}

func (k *KVStore) Set(key, value string) {
	k.mutex.Lock()
	defer k.mutex.Lock()
	k.store[key] = value
}

func (k *KVStore) Delete(key string) {
	k.mutex.Lock()
	defer k.mutex.Unlock()
	delete(k.store, key)
}

func main() {
	store := NewKVStore()

	store.Set("username", "avinash")
	store.Set("role", "engineer")

	if val, ok := store.Get("username"); ok {
		fmt.Println("✅ username:", val)
	}

	store.Delete("role")
	if _, ok := store.Get("role"); !ok {
		fmt.Println("❌ role deleted")
	}
}
