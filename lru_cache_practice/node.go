package main

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func NewNode(key, val int) *Node {
	return &Node{
		key:   key,
		value: val,
		next:  nil,
		prev:  nil,
	}
}
