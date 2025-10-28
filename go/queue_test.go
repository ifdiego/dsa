package queue

import "testing"

func TestQueue(t *testing.T) {
	q := Queue[int]{}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	val, _ := q.Peek()
	if val != 1 {
		t.Errorf("peek error, expected 1, got %v", val)
	}

	deq, _ := q.Dequeue()
	if deq != 1 {
		t.Errorf("dequeue error, expected 1, got %v", val)
	}

	if q.IsEmpty() {
		t.Errorf("error, queue shouldn't be empty")
	}
}
