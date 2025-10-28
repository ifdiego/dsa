package linkedlist

import "fmt"

type DoublyLinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Equals func(a, b T) bool
	Size   int
}

func (l *DoublyLinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Size++
		return
	}

	newNode.Previous = l.Tail
	if l.Tail != nil {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
	l.Size++
}

func (l *DoublyLinkedList[T]) TraverseForward() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
	fmt.Println("")
}

func (l *DoublyLinkedList[T]) TraverseBackward() {
	current := l.Tail
	for current != nil {
		fmt.Println(current.Value)
		current = current.Previous
	}
	fmt.Println("")
}

func (l *DoublyLinkedList[T]) Search(value T) *Node[T] {
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

func (l *DoublyLinkedList[T]) Delete(value T) {
	if l.Head == nil || l.Equals == nil {
		return
	}

	current := l.Head
	for current != nil {
		if l.Equals(current.Value, value) {
			if current.Previous == nil && current.Next == nil {
				l.Head = nil
				l.Tail = nil
			} else if current.Previous == nil {
				l.Head = current.Next
				l.Head.Previous = nil
			} else if current.Next == nil {
				l.Tail = current.Previous
				l.Tail.Next = nil
			} else {
				current.Previous.Next = current.Next
				current.Next.Previous = current.Previous
			}
			l.Size--
			return
		}
		current = current.Next
	}
}
