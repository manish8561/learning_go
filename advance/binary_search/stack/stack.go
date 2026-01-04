package stack

import "errors"

type StackInterface interface {
	Push(int)
	Pop() int
	IsEmpty()
}

// LIFO
type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Underflow")
	}
	l := len(s.items)
	ele := s.items[l-1]
	s.items = s.items[:l-1]
	return ele, nil
}
