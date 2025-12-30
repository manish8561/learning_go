package main

import (
	"fmt"
	"time"
)

// interview question
// closure function returning function and return value
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// Fibonacci as normal recursive function
func fibonacciNormal(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciNormal(n-1) + fibonacciNormal(n-2)
}

// Fibonacci function with memoization using a closure
func fibonacci() func(int) int {
	cache := map[int]int{}

	var fib func(n int) int // Declare the function variable for recursion
	fib = func(n int) int {
		if n <= 1 {
			return n
		}
		if result, ok := cache[n]; ok {
			return result
		}
		cache[n] = fib(n-1) + fib(n-2)
		return cache[n]
	}
	return fib
}

func main() {
	//returning function
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println("---------------------------------------")
	fib := fibonacci()
	//compare normal fibonacci and closure fibonacci with memoization
	tNow := time.Now()

	n := 50

	for i := range n {
		fmt.Printf("fibonacci Normal(%d) = %d\n", i, fibonacciNormal(i))

	}
	fmt.Println(time.Since(tNow).Microseconds())

	fmt.Println("---------------------------------------")
	//closure fibonacci with memoization
	tNow = time.Now()
	for i := range n {
		fmt.Printf("fibonacci with cache(%d) = %d\n", i, fib(i))
	}
	fmt.Println(time.Since(tNow).Microseconds())
}
