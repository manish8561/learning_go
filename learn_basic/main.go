package main

import "fmt"

// defining constants
const (
	first = iota + 2
	second
	third = 'a'
	fourth
	fivth = "mmmmmmm"
	sixth
)
// we can also use package reflect.TypeOf or %T in Printf
// interview question
// type switch example in golang
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case float32, float64:
		fmt.Printf("Float value:  %v type %T\n", v, v)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}


func main() {
	mmap := map[string]string{"0":"a", "1":"b", "2":"c"}

	for key,value := range mmap{
		fmt.Println("Key: ", key, "Value: ", value)
	}

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
		for j := 0; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)

	// example to test types with switch
	do(21)
	do(3.14)
	do("hello")
	do(true)
}

