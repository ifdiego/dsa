package stack

type stack[T any] struct {
	items []T
}

func (s *stack[T]) push(value T) {
	s.items = append(s.items, value)
}

func (s *stack[T]) pop() (T, bool) {
	if s.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *stack[T]) peek() (T, bool) {
	if s.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return s.items[len(s.items)-1], true
}

func (s *stack[T]) isEmpty() bool {
	return len(s.items) == 0
}

func (s *stack[T]) size() int {
	return len(s.items)
}
