package stack

import "testing"

func TestStackWithArray(t *testing.T) {
	stack := stack[int]{}

	stack.push(1)
	stack.push(2)
	stack.push(3)

	if stack.size() != 3 {
		t.Errorf("expect size of 3, but got %d", stack.size())
	}

	top, ok := stack.peek()
	if !ok || top != 3 {
		t.Errorf("expect top 3, but got %d", top)
	}

	pop, _ := stack.pop()
	if !ok || top != 3 {
		t.Errorf("expect pop 3, but got %d", pop)
	}
}
