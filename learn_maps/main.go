package main

import "fmt"

func main() {
	var map1 map[string]int

	fmt.Println(map1 == nil)

	map2 := map[int]string{112: "friend", 113: "neighbour"}

	map2[114] = "neighbour2"

	for k, v := range map2 {
		fmt.Println("K: ", k, "V: ", v)
	}
	// delete key from map
	delete(map2, 114)

	val, ok := map2[114]

	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("map2 after deletion: ", map2)
	}
	// use of make to create map
	map3 := make(map[string]int)

	map3["one"] = 1
	map3["two"] = 2

	// using %#v to print the map with its type
	fmt.Printf("using make function to create map map3: %#v", map3)
	fmt.Println()

	//more examples of map keys and values
	example()
}

func example() {
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
