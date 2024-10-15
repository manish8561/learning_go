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
}
