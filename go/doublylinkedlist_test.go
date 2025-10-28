package linkedlist

import "testing"

func TestDoublyLinkedList(t *testing.T) {
	list := DoublyLinkedList[int]{
		Equals: func(a, b int) bool { return a == b },
	}

	list.Append(42)
	list.Append(8)
	list.Append(99)
	list.Append(31)

	list.TraverseForward()
	list.TraverseBackward()

	list.Delete(8)
	list.TraverseForward()
	list.TraverseBackward()

	list.Delete(31)
	list.TraverseForward()
	list.TraverseBackward()
}
