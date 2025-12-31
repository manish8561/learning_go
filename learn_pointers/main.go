package main

import "fmt"

func main() {
	var a *int
	var b *int

	a = ptr(10)
	b = ptrG(20)

	c := *a + *b

	fmt.Println("memory address of a: ", &a)
	fmt.Println("memory address of b: ", &b)
	fmt.Println(c)

	d := new(int)
	fmt.Println("Pointer with new keyword: ", &d)

	//pointer to pointer
	var e *int
	var f **int
	e = ptr(30)
	f = &e

	fmt.Println("value of e: ", *e)
	fmt.Println("value of f: ", **f)
	fmt.Println("memory address of e: ", &e)
	fmt.Println("memory address of f: ", &f)
	fmt.Println("memory address stored in f: ", f)
	fmt.Println("value stored at the address stored in f: ", *f)
}

// generic after 1.18
func ptrG[T any](v T) *T {
	return &v
}

func ptr(v int) *int {
	return &v
}
