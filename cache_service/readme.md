# In-Memory Key-Value Store (Redis-like) - LLD Question

## Scenario
Design a simple in-memory key-value store similar to Redis. The store should support basic operations, optional expiry of keys, and safe concurrent access. The system will **not persist data to disk**, but must be **thread-safe** and support basic TTL (time-to-live) functionality.

---

## Requirements

### 1. Key-Value Operations
- `SET(key, value)` → store a key-value pair.
- `GET(key)` → retrieve value for a key.
- `DEL(key)` → delete a key from the store.
- `EXPIRE(key, seconds)` → set TTL for a key. After TTL expires, the key is automatically deleted.
- `TTL(key)` → returns remaining time-to-live for a key in seconds. If no TTL exists, return -1.

### 2. Concurrency
- Multiple goroutines may access the store concurrently.  
- Must ensure **thread-safe operations** on keys.

### 3. Optional Enhancements
- Support `INCR(key)` / `DECR(key)` for integer values.
- Support **batch retrieval**: `MGET(keys...)`.
- Simple eviction strategy (like LRU) if memory limit is reached (optional).

---

## Scope Limits
- **In-memory only** (no disk/database persistence).  
- **Single-machine** implementation (no distributed support).  
- TTL accuracy can be approximate (seconds-level granularity).  
- Focus on correctness, thread safety, and TTL mechanism.  
- Avoid over-engineering; no need for replication, clustering, or persistence.  

---

## Sample Usage
```go
store := NewKVStore()

store.SET("name", "Avinash")
value := store.GET("name") // returns "Avinash"

store.EXPIRE("name", 5) // key expires in 5 seconds
ttl := store.TTL("name") // returns remaining TTL in seconds

store.DEL("name")


lets think about the TTL little later. 



For expire, I thik I have a logic. 


I wont implement any logic to keep things deleting. 
I will put a expiry time for each key. 
and when we try to access that key, we compare that time with current time, if that time is past, then we delete the key and also return that we dont have anything. 
