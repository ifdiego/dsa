package linkedlist

import "fmt"

type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}

type LinkedList[T any] struct {
	Head   *Node[T]
	Equals func(a, b T) bool
	Size   int
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	l.Size++
}

func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: l.Head}
	l.Head = newNode
	l.Size++
}

func (l *LinkedList[T]) Traverse() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
	fmt.Println("")
}

func (l *LinkedList[T]) Search(value T) *Node[T] {
	if l.Head == nil || l.Equals == nil {
		return nil
	}

	current := l.Head
	for current != nil {
		if l.Equals(current.Value, value) {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *LinkedList[T]) Pop() (T, bool) {
	if l.Head == nil {
		var zeroValue T
		return zeroValue, false
	}

	if l.Head.Next == nil {
		value := l.Head.Value
		l.Head = nil
		l.Size--
		return value, true
	}

	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}

	value := current.Next.Value
	current.Next = nil
	l.Size--
	return value, true
}

func (l *LinkedList[T]) Dequeue() (T, bool) {
	if l.Head == nil {
		var zeroValue T
		return zeroValue, false
	}

	value := l.Head.Value
	l.Head = l.Head.Next
	l.Size--
	return value, true
}
