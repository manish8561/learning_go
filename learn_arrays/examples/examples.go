package examples

import (
	"fmt"
)

func removeElement(slice []int, element int) []int {
	for i, v := range slice {
		if v == element {
			// Remove the element by appending the slices before and after the index
			return append(slice[:i], slice[i+1:]...) // this way can also used to merge array
		}
	}
	return slice // Return the original slice if the element was not found
}

func Example1() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", numbers) // Output: [10 20 30 40 50]

	// Remove the element with the value 50
	numbers = removeElement(numbers, 50)

	fmt.Println("Modified slice:", numbers) // Output: [10 20 40 50]
}
