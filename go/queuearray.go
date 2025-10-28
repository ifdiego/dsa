package queue

type queue[T any] struct {
	items []T
}

func (q *queue[T]) enqueue(value T) {
	q.items = append(q.items, value)
}

func (q *queue[T]) dequeue() (T, bool) {
	if q.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return item, true
}

func (q *queue[T]) peek() (T, bool) {
	if q.isEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return q.items[len(q.items)-1], true
}

func (q *queue[T]) isEmpty() bool {
	return len(q.items) == 0
}

func (q *queue[T]) size() int {
	return len(q.items)
}
