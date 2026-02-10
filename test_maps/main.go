package main

import "fmt"

func main() {
	//map keys are values accessed not referenced, so they can be pointers, but they are not dereferenced when used as keys.
	var i *int
	i = new(int)
	*i = 42
	fmt.Println(*i)

	m := make(map[*int]interface{})
	fmt.Println(m[i])

	// n:=make(map[[]int]int) // invalid map key type []int
	// map[[int]string]string // invalid map key type [int]string
	// map[int]map[string]string // valid map key type int
}
