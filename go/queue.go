package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{
		value: value,
	}
	if q.tail == nil {
		q.head = newNode
		q.tail = newNode
		return
	}
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var zero T
		return zero, false
	}
	return q.head.value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func main() {
	queue := Queue[int]{}
	nums := []int{12, 6, 18, 19, 21, 11, 3, 5, 4, 24, 18}

	for _, num := range nums {
		queue.Enqueue(num)
	}

	deq, _ := queue.Dequeue()
	fmt.Println("dequeue:", deq)
	peek, _ := queue.Peek()
	fmt.Println("peek:", peek)
}
