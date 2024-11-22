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

func addLine() {
	fmt.Println("===================================================")
}

// we can also use package reflect.TypeOf or %T formatter in Printf function
// interview question
// type switch example in golang
// also example of polymorphism
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

	fmt.Println("const values", first, second, third, fourth, fivth, sixth)
	addLine()

	// hackerrank question to count digit as divisor for specific number
	// e.g. 123 like 1 & 3 is divisor for this number
	count := 0
	num := 123456789
	temp := num
	for temp > 0 {
		d := temp % 10
		// fmt.Println(num, " % ", d, num%d == 0)
		if d != 0 && num%d == 0 {
			count++
		}
		temp = temp / 10
	}

	fmt.Println(count)

	// map
	mmap := map[string]string{"0": "a", "1": "b", "2": "c"}

	for key, value := range mmap {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	addLine()

	arr := [...]int{3, 5, 6, 4, 8, 7, 9, 1}

	// bubble sort desecending order
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] { // > for ascending
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

	// insertion sort in ascending order
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
	addLine()

	// example to test types with switch
	do(21)
	do(3.14)
	do("hello")
	do(true)
	addLine()

	// new keyword example in golang
	newExample()
}

func newExample() {
	// diff b/w new and make in golang
	/*
	 * new() allocates memory and returns a pointer to the newly allocated zero value of the specified type.
	 * make() is used for initializing slices, maps, and channels, and it returns an initialized (but not zeroed) value of that type.
	 */

	// Using new() to allocate memory for an int
	numPtr := new(int) // numPtr is of type *int, pointing to a zero value (0)

	fmt.Println("Value of numPtr:", *numPtr, &numPtr) // Dereferencing to get the value (initially 0)

	// Updating the value through the pointer
	*numPtr = 42

	fmt.Println("Updated value of numPtr:", *numPtr) // Dereferencing to get the updated value (42)

	// Using new() to allocate memory for a custom struct
	type Person struct {
		name string
		age  int
	}

	personPtr := new(Person) // personPtr is of type *Person, pointing to a zeroed Person struct

	fmt.Println("Person:", *personPtr, "Memory Address: ", &personPtr) // Output: {"" 0}

	// Updating the struct through the pointer
	personPtr.name = "Alice"
	personPtr.age = 30

	fmt.Println("Updated Person:", *personPtr) // Output: {Alice 30}
}
