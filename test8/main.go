package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// interview questions
// goroutines
// channels
// types

// strings

func task1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Task 1 iterations: ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func task2() {
	for i := 0; i < 5; i++ {
		fmt.Println("Task 2 iterations: ", i)
		time.Sleep(50 * time.Millisecond)
	}
}

func main() {
	fmt.Println("-----------------------------------")

	// light weight? go keyword, go funciton()
	// channels for communication b/w goroutines, blocking in nature
	// types of channels buffered and unbuffered
	// diff in buffered and unbuffered channels
	// limit and unlimit
	// async read or write operations whereas sync reader or writer operations both must be present at the same time.

	// make() => maps, channels, []slices
	ch := make(chan int) // unbuffered

	// ch <- 2 // only sending value, output deadlock

	go func() {
		ch <- 2
	}()

	value := <-ch // only receive, output deadlock

	fmt.Println(value)

	fmt.Println("-----------------------------------")
	// concurrency and parallelism
	// memory and resources consumption is more in parallelism

	// concurrency example
	// go task1()
	// go task2()
	// time.Sleep(1 * time.Second)

	// parallelism example
	runtime.GOMAXPROCS(2) // allcating 2 cpu for goroutine

	go task1()
	go task2()
	time.Sleep(1 * time.Second) // not recommended in prod build

	//string example
	fmt.Println("-----------------------------------")

	input := "abc,xyz,poi, abc, poi,xyz"

	inputSlice := strings.Split(input, ",")

	if len(inputSlice) == 0 {
		fmt.Println("")
		return
	}

	checkValues := make(map[string]bool)

	for _, val := range inputSlice {
		val := strings.Trim(val, " ")
		if _, ok := checkValues[val]; ok {
			continue
		}
		checkValues[val] = true
	}
	// converting map to slice
	uniqueValues := make([]string, len(checkValues))
	counter := 0
	for keys := range checkValues {
		uniqueValues[counter] = keys
		counter++
	}
	fmt.Println(strings.Join(uniqueValues, ","))
}
