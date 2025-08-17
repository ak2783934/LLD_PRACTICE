package main

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

func NewNode(key, val int) *Node {
	return &Node{
		Key: key,
		Val: val,
	}
}
