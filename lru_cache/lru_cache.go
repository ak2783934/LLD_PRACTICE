package main

import (
	"fmt"
	"sync"
)

type LRUCache struct {
	Capacity int
	LRUMap   map[int]*Node
	List     DoublyLinkedList
	Len      int
	Mu       sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		LRUMap:   make(map[int]*Node),
		List:     *NewEmptyDoublyLinkedList(),
		Len:      0,
	}
}

func (l *LRUCache) GetValue(key int) int {
	// check in the maps
	// if no, return -1
	// if yet, return value and update the list

	l.Mu.Lock()
	defer l.Mu.Lock()
	node, ok := l.LRUMap[key]
	if !ok {
		return -1
	}

	l.List.RemoveNode(node)
	l.List.AddToFront(key, node.Val)
	l.LRUMap[key] = l.List.Tail
	return node.Val
}

func (l *LRUCache) SetValue(key, value int) {
	l.Mu.Lock()
	defer l.Mu.Lock()
	node, ok := l.LRUMap[key]
	if ok {
		// delete the existing node.
		l.List.RemoveNode(node)
	} else {
		l.Len++
	}

	// set the new values and update.
	l.List.AddToFront(key, value)
	l.LRUMap[key] = l.List.Tail
	if l.Len > l.Capacity {
		l.List.RemoveFromHead()
	}
}

func (l *LRUCache) PrintLRUCache() {
	if l.List.Head == nil {
		fmt.Println("Cache is empty")
		return
	}

	listHead := l.List.Head

	for listHead != nil {
		fmt.Printf(" key: %d val: %d", listHead.Key, listHead.Val)
		listHead = listHead.Next
	}
	fmt.Println()
}
