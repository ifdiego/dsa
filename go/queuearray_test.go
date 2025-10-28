package queue

import "testing"

func TestQueueWithArray(t *testing.T) {
	queue := queue[int]{}

	queue.enqueue(42)
	queue.enqueue(101)
	queue.enqueue(99)

	if queue.size() != 3 {
		t.Errorf("expect size 3, but got %v", queue.size())
	}

	top, ok := queue.peek()
	if !ok || top != 99 {
		t.Errorf("expect top 99, but got %d", top)
	}

	deq, _ := queue.dequeue()
	if !ok || deq != 42 {
		t.Errorf("expect dequeue 42, but got %d", deq)
	}
}
