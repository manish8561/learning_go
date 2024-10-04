package main

import "fmt"

// interview question
// generic function example in golang
func Sum[T int | float32 | float64](a, b T) T {
	return a + b
}

func main() {
	x, y := 2, 4
	c, d := 4.5, 5.6
	fmt.Println("int sum: ", Sum(x, y))
	fmt.Println("float sum:", Sum(c, d))
}
