package main

type LRUCache struct {
	size     int
	capacity int
	Queue    *DoublyLinkedList
	HashMap  map[int]*Node
}

func CreateLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		size:     0,
		capacity: capacity,
		Queue:    CreateDoublyLinkedList(),
		HashMap:  make(map[int]*Node),
	}
}

func (l *LRUCache) get(key int) int {
	// find if the key exist.

	// if not, then return -1

	// else pick value
	// delete the existing node.
	// push to head
	// update the map
	node, ok := l.HashMap[key]
	if !ok {
		return -1
	}
	value := node.value
	l.Queue.RemoveNode(node)
	l.Queue.AddToHead(node)

	return value
}

func (l *LRUCache) set(key, val int) {
	// find if the key exists
	// delete from map and queue
	// else
	// create a node
	// append and add into map and do more things.
	// if size increase limit, then remove from tail

	node, ok := l.HashMap[key]
	if ok {
		node.value = val
		l.Queue.RemoveNode(node)
		l.Queue.AddToHead(node)
		return
	} else {
		l.size++
	}

	newNode := NewNode(key, val)
	l.Queue.AddToHead(newNode)
	l.HashMap[key] = newNode

	if l.size > l.capacity {
		tail := l.Queue.RemoveTail()
		delete(l.HashMap, tail.key)
		l.size--
	}
}
