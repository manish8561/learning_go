package main

import "fmt"

// testing for benchmarking with test file check the main_test.go file
func SquareValue(v int) int{
	return v*v
}

func main(){
	res := SquareValue(2)
	fmt.Println("Result: ", res)
}