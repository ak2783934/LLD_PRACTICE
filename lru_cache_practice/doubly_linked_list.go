package main

import "fmt"

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func CreateDoublyLinkedList() *DoublyLinkedList {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &DoublyLinkedList{Head: head, Tail: tail}
}

func (d *DoublyLinkedList) AddToHead(node *Node) {
	next := d.Head.next
	d.Head.next = node
	node.prev = d.Head
	node.next = next
	next.prev = node
}

func (d *DoublyLinkedList) RemoveNode(node *Node) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	node.next = nil
	node.prev = nil
}

func (d *DoublyLinkedList) RemoveTail() *Node {
	if d.Tail == d.Head {
		return nil
	}

	tailNode := d.Tail.prev
	d.RemoveNode(tailNode)
	return tailNode
}

func (d *DoublyLinkedList) PrintList() {
	curr := d.Head.next

	for curr != nil && curr != d.Tail {
		fmt.Print(curr.key, ":", curr.value, " -> ")
		curr = curr.next
	}
	fmt.Println("NULL")
}
