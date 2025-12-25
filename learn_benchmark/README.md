# Go Benchmarking Learning

This folder demonstrates how to write and run benchmarks in Go.

## Overview

Benchmarking is a way to measure the performance of your code. Go provides built-in support for benchmarking through the `testing` package.

## Files

- **main.go**: Contains the `SquareValue()` function that calculates the square of an integer
- **main_test.go**: Contains the benchmark test for performance measurement

## Code Structure

### main.go

```go
func SquareValue(v int) int {
    return v*v
}
```

This is a simple function that returns the square of a given integer value.

### main_test.go

```go
func BenchmarkSquareValue(b *testing.B) {
    for i := 0; i < b.N; i++ {
        SquareValue(i)
    }
}
```

The benchmark function follows the naming convention `Benchmark[FunctionName]` and accepts a `*testing.B` parameter.

## Running Benchmarks

### Run benchmark tests:

```bash
go test -bench=. -count=10
```

- `-bench=.` : Run all benchmarks
- `-count=10` : Run each benchmark 10 times for more reliable results

### Check test coverage:

```bash
go test -cover
```

### Run the application:

```bash
go run main.go
```

Output: `Result: 4` (2 squared)

## Key Concepts

1. **Benchmark Function**: Must start with `Benchmark` and accept `*testing.B`
2. **b.N**: The number of iterations the benchmark loop will execute
3. **Performance Measurement**: Go automatically measures execution time and memory allocation
4. **Multiple Runs**: Using `-count` flag helps identify consistent performance patterns

## Learning Outcomes

- Understanding how to write benchmark tests in Go
- Measuring function performance
- Identifying performance bottlenecks
- Using `-bench` and `-cover` flags with `go test`
