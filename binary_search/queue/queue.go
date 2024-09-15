package queue

import "errors"

// type QueueInterface interface {
// 	Enqueue(int)
// 	Dequeue() (int, error)
// 	Peek() (int, error)
// 	IsEmpty() bool
// }
// FIFO
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Underflow")
	}
	ele := q.items[0]
	q.items = q.items[1:]

	return ele, nil
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Underflow")
	}
	return q.items[0], nil
}
