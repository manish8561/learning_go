package main

import (
	"fmt"
	"sync"
	"time"
)

// panic and recover example
var wg sync.WaitGroup

func cleanup() {
	defer wg.Done()

	if r := recover(); r != nil {
		fmt.Println("Recoverd in cleanup", r)
	}
}

func say(s string) {
	defer cleanup()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("On dear, a 2")
		}
	}
}

func testPanic() {
	panic("testing panic")
	// fmt.Println("This line will never execute")
}

func otherRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	testPanic()
	fmt.Println("If panic occur this will never execute")
}

func main() {
	// example of simple panic and recover
	otherRecover()

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Add(1)
	go say("Hi")
	wg.Wait()

	time.Sleep(time.Second)

}
