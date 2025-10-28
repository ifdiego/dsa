package linkedlist

import "testing"

func TestLinkedList(t *testing.T) {
	list := LinkedList[int]{}

	list.Append(1)
	list.Append(2)
	list.Append(3)

	list.Traverse()

	value, ok := list.Pop()
	if !ok || value != 3 {
		t.Errorf("pop failed, expect 3, got %v", value)
	}

	deq, ok := list.Dequeue()
	if !ok || deq != 1 {
		t.Errorf("dequeue failed, expected 1, got %v", deq)
	}
}
