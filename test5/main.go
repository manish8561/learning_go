package main

import "fmt"

// interview question
// simple example with defer, panic and recover
// calling multiple defer in the stack
func main() {
	// defer will execute even with errors
	defer fmt.Println("testing defer call as stack")
	defer func() {
		// recover from panic
		if r := recover(); r != nil {
			fmt.Println("some panic error in the funciton")
		}
	}()
	defer fmt.Println("testing defer call as stack 2")

	panic("something went wrong")

	// unreachale code
	fmt.Println("hihihihihi")
}
