package main

import (
	"testing"
)

// command to run bench testing
// go test -bench=. -count=10
// command for coverage in test
// go test -cover
func BenchmarkSquareValue(b *testing.B) {
	n := b.N
	b.Log("value of bench N: ", n)
	for i := 0; i < n; i++ {
		// Code to be benchmarked
		SquareValue(i)
	}
}

func BenchmarkPrintValue(b *testing.B) {
	n := b.N
	b.Log("value of bench N: ", n)
	PrintValue(n)
}
