package examples

import "fmt"
// tricky interveiw questions of slices 
func Example1() {
	something := make([]int, 4) // slice

	b := something[0:1]

	b[0] = 1000 // since value in slice are called by reference

	fmt.Println(b, something)
}

func Example2() {
	something := make([]int, 4) // slice

	something = append(something, 1) // capacity become double if not specified
	fmt.Println(something, len(something), cap(something))
}

func GetValues(a []int) { // value is changed in case of slices
	a[0] = 20
}

func Example3() {
	something := make([]int, 4, 10) // slice

	something = append(something, 1)
	fmt.Println(something, len(something), cap(something))

	GetValues(something) // pass by reference change the value
	// something[7] = 1000 // this line gives error
	fmt.Println(something)

}
