package main

import "fmt"

func main(){
	var a *int
	var b *int

	a = ptr(10)
	b = ptrG(20)

	c := *a + *b

	fmt.Println("memory address of a: ", &a)
	fmt.Println("memory address of b: ", &b)
	fmt.Println(c)

}
//generic after 1.18
func ptrG[T any](v T) *T{
	return &v
}



func ptr(v int) *int{
	return &v
}