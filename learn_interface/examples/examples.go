package examples

import "fmt"

type values struct {
	first  int
	second int
}

type mathsTest interface {
	add(a, b int) int
	sub(a, b int) int
	multiply(a, b int) int
	// divide(a, b int) int
}

func (v values) add(a, b int) int {
	return a + b + v.first + v.second
}

func (v values) multiply(a, b int) int {
	return a * b * v.first * v.second
}

func (v values) sub(a, b int) int {
	return a - b - v.first - v.second
}

func Example() {
	var t mathsTest = values{1, 2}
	fmt.Println(t.add(1, 2))
	fmt.Println(t.multiply(1, 2))
	fmt.Println(t.sub(1,2))
}
