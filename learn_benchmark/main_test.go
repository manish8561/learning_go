package main

import (
    "testing"
)
// command to run bench testing
// go test -bench=. -count=10
// command for coverage in test
// go test -cover
func BenchmarkSquareValue(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Code to be benchmarked
		SquareValue(i)
    }
}