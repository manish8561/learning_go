package main

import "fmt"

// interview question
// simple example with defer, panic and recover
func main() {
	defer func ()  {
		// recover from panic
		if r := recover(); r != nil {
			fmt.Println("some panic error in the funciton")
		}
	}()
	panic("something went wrong")
	fmt.Println("hihihihihi")
}