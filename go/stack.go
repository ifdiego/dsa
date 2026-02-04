package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	list *Node[T]
}

func (s *Stack[T]) Push(value T) {
	newNode := &Node[T]{
		value: value,
		next:  s.list,
	}
	s.list = newNode
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.list == nil {
		var zero T
		return zero, false
	}
	value := s.list.value
	s.list = s.list.next
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.list == nil {
		var zero T
		return zero, false
	}
	return s.list.value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.list == nil
}

func main() {
	stack := Stack[int]{}
	nums := []int{12, 6, 18, 19, 21, 11, 3, 5, 4, 24, 18}

	for _, num := range nums {
		stack.Push(num)
	}

	pop, _ := stack.Pop()
	fmt.Println("pop:", pop)
	peek, _ := stack.Peek()
	fmt.Println("peek:", peek)
}
