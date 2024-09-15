package examples

import (
	"fmt"
	"sync"
	"time"
)

/* concurrency in golang */

// example using wait group
func count(thing string) {
	for i := 1; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 100)
	}
}
func Example1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		count("sheep")
		wg.Done()
	}()
	wg.Wait()
}

// example using channels without wait group
func count2(thing string, c chan string) {
	for i := 1; i < 10; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func Example2() {
	c := make(chan string)
	go count2("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

// example using buffered channel
func Example3() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}

// example using select
func Example4() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func Example5() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 0; i < 100; i++ {
		fmt.Println(<-results)
	}

}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}

}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
