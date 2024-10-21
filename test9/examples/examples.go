package examples

import "fmt"

/**
 * car in circular road with juice machines on the road half the length of the array
 * and pairs fuel in the cart with distance need to travel (1 Juice/ distance )
 * give me first occurrence of the pair that will complete the circule distance successfully.
 * 
 * 8
 * 2 3 4 5 6 7 4 6
 * 6
 * 2, 3, 6, 5, 8, 7
 */

func Example1() int {
	arr := [...]int{2, 3, 6, 5, 8, 7, 4, 6}

	count := -1
	for i := 0; i < len(arr)-1; {
		count++
		fmt.Println(arr[i], " >= ", arr[i+1])
		if arr[i] >= arr[i+1] {
			return count
		}

		i += 2
	}
	return -1
}
