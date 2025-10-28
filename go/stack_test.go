package stack

import "testing"

func TestStack(t *testing.T) {
	s := Stack[int]{}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	peek, _ := s.Peek()
	if peek != 3 {
		t.Errorf("peek error, expected 3, got %v", peek)
	}

	val, _ := s.Pop()
	if val != 3 {
		t.Errorf("pop error, expected 3, got %v", val)
	}

	if s.IsEmpty() {
		t.Errorf("error, stack shouldn't be empty")
	}
}
