package main

import (
	"fmt"

	calculator "example.com/demo/packages"
)

func init() {
	//testing public variables
	calculator.Capital_var = "manish"
}

func main() {

	fmt.Println("Hello World!", calculator.Capital_var)
	a, b, c := 5, 3, 0

	c = calculator.Add(a, b)
	fmt.Println("Result", c)
	c = calculator.Sub(a, b)
	fmt.Println("Result", c)

	s1 := make([]int, 10, 20)
	fmt.Println(s1)

	s2 := make(map[string]int)
	fmt.Println(s2)

	s3 := make(chan int)
	fmt.Println(s3)

}
