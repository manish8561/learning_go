package main

import "fmt"

// interview question
/**
 * question to get counts of repeated elements from the array and also get the maximum value
 */
func main() {
	arr := [...]int{1, 2, 2, 2, 3, 3, 4, 4, 1, 2, 3}

	m := make(map[int]int)
	max := 0

	for _, v := range arr {
		// check value in golang map exists or not
		if val, ok := m[v]; ok {
			// val++
			// m[arr[i]] = val
			m[v] = val + 1
		} else {
			m[v] = 1
		}

		if max < m[v] {
			max = m[v]
		}
	}

	fmt.Println(m, max)
}
