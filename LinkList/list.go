package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	position int
	head     *Node
}

func (l *List) enqueue() {
	n := &Node{value: l.position}
	current := l.head
	l.position++
	if l.head != nil {
		for current.next != nil {
			current = current.next
		}
		current.next = n
		return
	}
	l.head = n
	return
}

func (l *List) dequeue() {
	l.head = l.head.next
}

func (l *List) print() {
	current := &l.head
	for current != nil {
		fmt.Println((*current).value)
	}
}

func main() {
	var list *List
	list = &List{
		head:     nil,
		position: 0,
	}

	list.enqueue()
	list.print()
	list.enqueue()
	list.print()
	list.enqueue()
	list.print()
	list.enqueue()
	list.print()
	list.dequeue()
	list.print()
	list.dequeue()
	list.print()
}
