package main

import "fmt"

// interview question
// generic function example in golang
func Sum[T int | float32 | float64](a, b T) T {
	return a + b
}

// example of generic type in golang as stack
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		panic("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	x, y := 2, 4
	c, d := 4.5, 5.6
	fmt.Println("int sum: ", Sum(x, y))
	fmt.Println("float sum:", Sum(c, d))

	fmt.Println("------------------------------")
	// generic type example
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)
	intStack.Push(4)
	fmt.Println(intStack.Pop())

	stringStack := Stack[string]{}
	stringStack.Push("Go")
	stringStack.Push("Generics")
	fmt.Println(stringStack.Pop())
}
