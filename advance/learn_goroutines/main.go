package main

import (

	"example.com/demo/examples"
)

/* concurrency in golang */
// example using worker groups with channels

func main() {
	// example 1
	examples.Example1()

	// example 2
	examples.Example2()

	// example 3
	examples.Example3()

	// example 4
	go examples.Example4()

	// example 5
	examples.Example5()
}