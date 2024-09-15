package main

import "fmt"

/* test for substring find in the string */
func check(L, N string) bool {
	M := make(map[int]int)
	charCount := 0
	for i := 0; i < len(L); i++ {
		charCode := int(L[i])
		if M[charCode] == 0 {
			charCount += 1
		}
		M[charCode] += 1
	}

	for i := 0; i < len(N); i++ {
		charCode := int(N[i])
		if M[charCode] > 0 {
			M[charCode] -= 1
			if M[charCode] == 0 {
				charCount -= 1
			}
		}

		if charCount == 0 {
			return true
		}
	}
	return false
}

// function to get first index of substring from a string
func find(str string, target string) int {
	for i := 0; i <= len(str)-len(target); i++ {
		found := true
		for j := 0; j < len(target); j++ {
			if str[i+j] != target[j] {
				found = false
			}
		}
		if found {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Printf("L string is sub string of N string")

	// calling check
	r := check("Manish", "Manish sharma")
	fmt.Println("Result:", r)

	//calling find
	res := find("Manish Sharma", "Shar")
	fmt.Println("Result 2:", res)
}
