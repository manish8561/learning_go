package main

import (
	"fmt"

	"example.com/u/gomonkeytesting/examples"
)

func main() {
	fmt.Println("Befor function call.....")
	r := examples.Add(5, 5)
	fmt.Println("After function call......... result: ", r)
}
