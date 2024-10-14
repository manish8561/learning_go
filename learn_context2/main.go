package main

import (
	"context"
	"fmt"
	"time"
)
// another feature of context
func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "123434")
}

func doSomethingCool(ctx context.Context) {
	rID := ctx.Value("request-id")
	fmt.Println(rID)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("time out")
		default:
			fmt.Println("doing something cool")
		}
		time.Sleep(500 * time.Millisecond)
	}

}
func main() {
	fmt.Println("Go Context Tutorial")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	ctx = enrichContext(ctx)
	go doSomethingCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("Oh no I have exceed the deadline")
		fmt.Println(ctx.Err())
	}
	time.Sleep(2 * time.Second)
}
