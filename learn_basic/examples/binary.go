package examples

// convert decimal to binary and vice versa

import (
	"fmt"
	"math"
	"strconv"
)

// DecimalToBinary converts a decimal number to its binary representation as a string.
func DecimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	var binaryStr string

	for n > 0 {
		remainder := n % 2
		// Prepend the remainder to the string to get the correct order(reading remainders bottom-up)
		binaryStr = strconv.Itoa(remainder) + binaryStr
		n /= 2
	}
	return binaryStr
}

// BinaryToDecimal converts a binary string to its decimal representation as an integer.
func BinaryToDecimal(b string) int {
	var decimal int
	length := len(b)

	for i := length - 1; i >= 0; i-- {
		position := length - 1 - i
		if b[i] == '1' {
			decimal += int(math.Pow(2, float64(position)))
		}
	}
	return decimal
}

// Example usage
func Example() {
	decimalNumber := 43
	binaryString := DecimalToBinary(decimalNumber)
	fmt.Printf("Decimal: %d -> Binary: %s\n", decimalNumber, binaryString)

	binaryInput := "101011"
	decimalValue := BinaryToDecimal(binaryInput)
	fmt.Printf("Binary: %s -> Decimal: %d\n", binaryInput, decimalValue)
}
