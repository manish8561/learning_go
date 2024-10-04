package main

import (
    "context"
    "fmt"
    "time"
)

func main() {
    // Create a context with a timeout of 2 seconds
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel() // Ensure that the context is canceled to release resources

    result := make(chan string)

    // Simulate some work in a goroutine
    go func() {
        time.Sleep(3 * time.Second) // Simulating a long-running task
        result <- "Task completed!"
    }()

    // Listen for the result or the context cancellation
    select {
    case res := <-result:
        fmt.Println(res)
    case <-ctx.Done():
        fmt.Println("Context timed out:", ctx.Err()) // context deadline exceeded
    }
}
