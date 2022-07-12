package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Queue struct {
	position int
	head     *Node
}

func (q *Queue) enqueue() {
	n := &Node{value: q.position}
	current := q.head
	q.position++
	if q.head != nil {
		for current.next != nil {
			current = current.next
		}
		current.next = n
		return
	}
	q.head = n
	return
}

func (q *Queue) dequeue() {
	q.head = q.head.next
}

func (q *Queue) print() {
	current := &q.head
	for current != nil {
		fmt.Println((*current).value)
	}
}

func main() {
	var queue *Queue
	queue = &Queue{
		head:     nil,
		position: 0,
	}

	queue.enqueue()
	queue.print()
	queue.enqueue()
	queue.print()
	queue.enqueue()
	queue.print()
	queue.enqueue()
	queue.print()
	queue.dequeue()
	queue.print()
	queue.dequeue()
	queue.print()
}
