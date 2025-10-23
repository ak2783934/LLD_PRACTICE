# LRU Cache — Machine Coding Problem

## 🎯 Objective

Implement a **Least Recently Used (LRU) Cache** that supports O(1) `get` and `put` operations using a Doubly Linked List (DLL) and a HashMap.

---

## 🔹 Functional Requirements

* Store **key-value** pairs.
* When capacity is reached, **evict the least recently used (LRU)** item.
* Every `get` or `put` operation should make the accessed key the **most recently used (MRU)**.

---

## 🔹 APIs to Implement (Go)

```go
type LRUCache struct {
    // your internal fields
}

func Constructor(capacity int) LRUCache
func (c *LRUCache) Get(key int) int
func (c *LRUCache) Put(key int, value int)
```

---

## 🔹 Constraints & Edge Cases

* Both `Get` and `Put` must run in **O(1)** time.
* Handle these edge cases explicitly:

  * `capacity = 0` — cache should store no items.
  * **Repeated keys** — `Put` on an existing key should update the value and move it to MRU.
  * **Eviction order** — remove the **tail** (LRU) on overflow.
  * Negative or zero keys/values are allowed.

---

## 🔹 Internal Behavior

* `Get(key)`: If key exists, return value and mark node as MRU; otherwise return `-1`.
* `Put(key, value)`: If key exists, update value and mark MRU. If key doesn't exist, insert as MRU. If capacity exceeded, evict LRU.

---

## 🔹 Example

```
LRUCache cache = Constructor(2);
cache.Put(1, 1);
cache.Put(2, 2);
cache.Get(1);       // returns 1 (1 becomes MRU)
cache.Put(3, 3);    // evicts key 2 (as 2 was LRU)
cache.Get(2);       // returns -1 (not found)
cache.Put(4, 4);    // evicts key 1
cache.Get(1);       // returns -1
cache.Get(3);       // returns 3
cache.Get(4);       // returns 4
```

---

## ✅ Interview Goals

1. Provide **clean, idiomatic Go** code.
2. Show **DLL + HashMap** usage clearly (node struct, head/tail sentinels).
3. Handle all edge cases and document decisions.
4. Discuss potential extensions: thread-safety, TTL, persistence.

---

When you finish, paste your full implementation here and I'll review for correctness, performance, and style.

assumptions: 
- key and values are always int. 
- 

// will try to keep node as premitive data type, since no major operations are going to happen on Node object. 
Node{
    key int
    value int
    next *Node
    prev *Node
}

func NewNode(key, value int) *Node

// DoublyLinkedList Head and Tail 
DoublyLinkedList{
    Head *Node
    Tail *Node
}

func CreateLinkedList() *DoublyLinkedList
func (d *DoublyLinkedList)AddToHead(node *Node) {}
func (d *DoublyLinkedList)RemoveTail() {}
func (d *DoublyLinkedList)RemoveNode(node *Node) {}


LRUCache {
    size int
    capacity int
    Queue *DoublyLinkedList
    HashMap map[int]*Node
}

func CreateLRUCache(int capacity) *LRUCache
func (l *LRUCache) get(key int)
func (l *LRUCache) set(key, value int)


