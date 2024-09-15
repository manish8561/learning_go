package main

import (
	"fmt"
	s "strings"
)

func main() {
	/* Need to import strings as s */
	fmt.Println(s.Contains("test", "e"))

	/* Build in */
	fmt.Println(len("hello")) // => 5
	// Outputs: 101
	fmt.Println("hello"[1], string("hello"[1]))
	// Outputs: e
	b := "hello2"
	fmt.Println(b[1])

	/*
		Contains("test", "es")	true
		Count("test", "t")	2
		HasPrefix("test", "te")	true
		HasSuffix("test", "st")	true
		Index("test", "e")	1
		Join([]string{"a", "b"}, "-")	a-b
		Repeat("a", 5)	aaaaa
		Replace("foo", "o", "0", -1)	f00
		Replace("foo", "o", "0", 1)	f0o
		Split("a-b-c-d-e", "-")	[a b c d e]
		ToLower("TEST")	test
		ToUpper("test")	TEST
	*/
	// this part is pending

	// naked returns
	fmt.Println("------------------------------")
	fmt.Println("Naked returns")
	x, y := split(17)
	fmt.Println(x) // => 7
	fmt.Println(y) // => 10

	// variadic functions works like functions overloading
	fmt.Println("------------------------------")
	fmt.Println("variadic functions")

	sum(1, 2)    //=> [1 2] 3
	sum(1, 2, 3) // => [1 2 3] 6

	nums := []int{1, 2, 3, 4}
	sum(nums...) // => [1 2 3 4] 10

	// function as values example
	fmt.Println("------------------------------")
	fmt.Println("functions as values")
	// assign a function to a name
	add := func(a, b int) int {
		return a + b
	}
	// use the name to call the function
	fmt.Println(add(3, 4)) // => 7

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
