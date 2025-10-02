package queue

import (
	"fmt"
)

// interview question

//	type QueueInterface interface {
//		Enqueue(int)
//		Dequeue() (int, error)
//		Peek() (int, error)
//		IsEmpty() bool
//	}
//
// FIFO
type Queue struct {
	items []int
}

// insert
func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

// Isempty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// delete
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("underflow")
	}
	ele := q.items[0]
	q.items = q.items[1:]

	return ele, nil
}

// peek
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, fmt.Errorf("underflow")
	}
	return q.items[0], nil
}
