package main

import "fmt"

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
func fibonacciNor(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciNor(n-1) + fibonacciNor(n-2)
}

// Fibonacci function with memoization using a closure
func fibonacci() func(int) int {
	cache := map[int]int{}
	var fib func(n int) int
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
	fmt.Println(fib(10)) // Output: 55
	fmt.Println(fib(20)) // Output: 6765

	// nth value from fibonacci
	fmt.Println(fibonacciNor(8))
}
