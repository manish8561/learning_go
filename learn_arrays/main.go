package main

import (
	"fmt"

	"example.com/arrays/examples"
)

func main() {
	// Declare and initialize an array
	var arr = [5]int{1, 2, 3, 4, 5}

	// Access and modify elements
	arr[0] = 10
	fmt.Println("First element:", arr[0])

	// Get the length of the array
	fmt.Println("Array length:", len(arr))

	// Iterate over the array using a for loop
	fmt.Println("Array elements:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// Iterate over the array using a range loop
	fmt.Println("Array elements with index:")
	for index, v := range arr {
		fmt.Println(index, v)
	}

	// Multidimensional array
	var matrix [3][3]int
	matrix[0][0] = 1
	matrix[1][1] = 2
	matrix[2][2] = 3
	fmt.Println("Matrix:", matrix)

	// Pass array to a function
	fmt.Println("Array before modification:", arr)
	modifyArray(&arr)
	fmt.Println("Array after modification:", arr)

	// comparing the arrays
	var a1 = [3]int{5, 6, 7}
	a2 := [...]int{5, 6, 7}
	var a3 = [3]int{6, 7, 8}

	fmt.Println(a1 == a2)
	fmt.Println(a1 == a3)
	fmt.Println(a2 == a3)

	// slices examples
	sliceExamples()
}

// Function to modify array
func modifyArray(arr *[5]int) {
	arr[0] = 100
}

// slices examples
func sliceExamples() {
	fmt.Println("---------------------------------------------")

	// array
	arr := [...]string{"This", "is", "the", "tutorial", "of", "Go", "slices"}

	// Display array
	fmt.Println(arr)

	myslice := arr[1:6]

	// print slices
	fmt.Println(myslice, len(myslice), cap(myslice))

	myslice = append(myslice, "testing")

	// print slices
	fmt.Println(myslice, len(myslice), cap(myslice))

	myslice = append(myslice, "testing")

	// print slices
	fmt.Println(myslice, len(myslice), cap(myslice))

	// loop over slices
	arr2 := []int{1, 2, 3, 4, 5}

	for index, value := range arr2 {
		fmt.Println("Index: ", index, "Value: ", value)
	}

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	slc := arr3[:4]

	fmt.Println("arr", arr3, "slc", slc)
	// values reassign will change value in array also since slices are reference types
	slc[0] = 10
	slc[1] = 20
	slc[2] = 100

	fmt.Println("arr", arr3, "slc", slc)

	// Another example
	fmt.Println("---------------------------------------------")
	examples.Example1()

	// testing condition for error
	arr4 := []int{}
	fmt.Println(arr4, "condition: ", arr4 == nil, "empty array: ", len(arr4) == 0)

}
