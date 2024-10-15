package main

import "fmt"

// interview question
// simple example with defer, panic and recover
// calling multiple defer in the stack
func main() {
	defer fmt.Println("testing defer call as stack")
	defer func ()  {
		// recover from panic
		if r := recover(); r != nil {
			fmt.Println("some panic error in the funciton")
		}
	}()
	panic("something went wrong")

	// unreachale code
	fmt.Println("hihihihihi")
}