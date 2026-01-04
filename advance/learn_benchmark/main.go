package main

import "fmt"

// testing for benchmarking with test file check the main_test.go file
func SquareValue(v int) int {
	return v * v
}

func PrintValue(v any) {
	fmt.Println("Value is: ", v)
}

func main() {
	res := SquareValue(2)
	PrintValue(res)
}
