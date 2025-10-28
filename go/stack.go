package stack

import (
	linkedlist "github.com/ifdiego/algorithms/linked-list"
)

type Stack[T any] struct {
	list linkedlist.LinkedList[T]
}

func (s *Stack[T]) Push(value T) {
	s.list.Append(value)
}

func (s *Stack[T]) Pop() (T, bool) {
	return s.list.Pop()
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.list.Head == nil {
		var zeroValue T
		return zeroValue, false
	}

	current := s.list.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list.Size == 0
}

func (s *Stack[T]) Print() {
	s.list.Traverse()
}
