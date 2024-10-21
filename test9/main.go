package main

import (
	"fmt"

	"example.com/u/demo/examples"
)

// interview question
// Print the prime number series for N th length

// Function to check if a number is prime
// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// Check for factors from 2 to n-1
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	// Example number to check
	n := 30
	for number := 1; number <= n; number++ {
		if isPrime(number) {
			fmt.Printf(" %d", number)
		}
	}
	fmt.Println()

	fmt.Println("-----------------------------------------")
	// car circular road return first pair to be completed
	r := examples.Example1()
	fmt.Println("result: ", r)
}
