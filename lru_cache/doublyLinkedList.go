package main

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewEmptyDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (d *DoublyLinkedList) AddToFront(key, val int) {
	node := NewNode(key, val)

	if d.Head == nil {
		d.Head = node
		d.Tail = node
		return
	}

	if d.Tail == nil {
		d.Tail = node
		d.Head.Next = d.Tail
		d.Tail.Prev = d.Head
		return
	}

	d.Tail.Next = node
	node.Prev = d.Tail
	d.Tail = node
}

func (d *DoublyLinkedList) RemoveFromHead() {
	d.Head = d.Head.Next
}

func (d *DoublyLinkedList) RemoveNode(node *Node) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		// node is head
		d.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		// node is tail
		d.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil
}
