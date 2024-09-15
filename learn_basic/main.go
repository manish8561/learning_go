package main

import "fmt"

const (
	first = iota + 2
	second
	third = 'a'
	fourth
	fivth = "mmmmmmm"
	sixth
)

func main() {
	fmt.Println("const values", first, second, third, fourth, fivth, sixth)

	arr := [...]int{3, 5, 6, 4, 8, 7, 9, 1}

	// bubble sort
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

	// insertion sort in desending order
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

/*
https://www.youtube.com/watch?v=lwKPqllpYCQ

*/