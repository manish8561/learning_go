package main

import "fmt"

func main() {
	even := []int{2,4,6,8,10,12,14,16,18,20}
	odd := []int{1,3,5,7,9,11,13,15,17,19}

	lt10 := even[0:5]
	lt10 = append(lt10, odd[0:5]...)

	fmt.Println(lt10)
	fmt.Println(even)// value changed due to call by reference slices
	fmt.Println(odd)
}