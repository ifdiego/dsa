package queue

import (
	linkedlist "github.com/ifdiego/algorithms/linked-list"
)

type Queue[T any] struct {
	list linkedlist.LinkedList[T]
}

func (q *Queue[T]) Push(value T) {
	q.list.Append(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.list.Dequeue()
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.list.Head == nil {
		var zeroValue T
		return zeroValue, false
	}
	return q.list.Head.Value, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.Size == 0
}

func (q *Queue[T]) Print() {
	q.list.Traverse()
}
