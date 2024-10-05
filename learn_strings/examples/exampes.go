package examples

import "fmt"

// interview question
/*
 * find out the odd from the series of string
 * ["ABC","BCD","CDE","DFI"]
 * length of above series is also given 4
 * difference in strings is (+1,+1), but in last string is (+2,+1), so output is
 * DFI
 */

// Helper function to check if the string has consecutive characters by +1
// func isConsecutive(s string) bool {
// 	for i := 0; i < len(s)-1; i++ {
// 		if s[i+1] != s[i]+1 {
// 			return false
// 		}
// 	}
// 	return true
// }

// Helper function to check if the string has a consistent step pattern
func hasStepPattern(s string) bool {
	if len(s) < 2 {
		return true
	}
	step := int(s[1]) - int(s[0])
	for i := 1; i < len(s)-1; i++ {
		if int(s[i+1])-int(s[i]) != step {
			return false
		}
	}
	return true
}

// Function to find the odd string in the series
func findOdd(series []string) string {
	for _, str := range series {
		if !hasStepPattern(str) {
			return str
		}
	}
	return ""
}

func Example1() {
	fmt.Println("-----------------------------------")
	// series := []string{"ABCDE", "DFIMR", "BCDEF", "CDEFG", "DEFGH"}
	series := []string{"ACE", "BDF", "CEG", "DGF"}

	r := findOdd(series)

	fmt.Println("Result: ", r)
}

// interview question
// strings array to be sorted in ascending order
// bubble sort
func sortString(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}

func Example2() {
	fmt.Println("---------------------------------")

	arr := []string{"Manish", "Sharma", "Ram", "Abhi", "Rahul", "Gaurav", "Vikram"}

	fmt.Println("Sorted arr of the string: ", sortString(arr))
}
