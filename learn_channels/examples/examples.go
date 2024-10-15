package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func DisplayLine() {
	fmt.Println("---------------------------------------------------------------------")
}

// Custom type for channel values
type Message struct {
	Sender    string
	Content   string
	Timestamp int64
}

// channel with custom type
func Example1() {
	defer DisplayLine()
	// Creating a channel that carries values of type 'Message'
	messageChannel := make(chan Message)

	// go routine receving value from channel
	go func() {
		// Sending the message through the channel
		// Receiving the value from the channel and status of open and close of channel
		receivedMsg, ok := <-messageChannel

		// Printing the received message
		fmt.Printf("Received message from %s: %s\n", receivedMsg.Sender, receivedMsg.Content)
		fmt.Println("channel status: ", ok)
	}()
	
	// Creating a Message instance
	msg := Message{
		Sender:    "Alice",
		Content:   "Hello, Bob!",
		Timestamp: 1620725555,
	}
	// Example of sending a value through the channel
	messageChannel <- msg
	close(messageChannel)

}

// channels with select statements and timeout in select
func Example2() {
	defer DisplayLine()

	fmt.Println("Example 2 with select statements")
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// example with wait group and atomic perform low-level call
/*
 * Use Cases
 * Safe access to shared variables
 * Performance optimization
 * Lock-free synchronization
 */
func Example3() {
	defer DisplayLine()

	fmt.Println("Example 3 for atomic counters and wait group")
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops.Load())
}

// interview question
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

// example with wait group and mutex
func Example4() {
	defer DisplayLine()

	fmt.Println("Example 4 wait group and mutex")
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}

// example of atomic without mutex and stateful go routines
func Example5() {
	defer DisplayLine()

	fmt.Println("Example 5 with stateful go routine")
	type readOp struct {
		key  int
		resp chan int
	}
	type writeOp struct {
		key  int
		val  int
		resp chan bool
	}

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

}

// loops with channels for receiving and check channel is closed or not
func Example6() {
	defer DisplayLine()

	fmt.Println("Example 6 for basic example and close channel")
	c := make(chan string)

	go initStrings(c)

	for {
		res, ok := <-c
		if !ok {
			fmt.Println("Channel is closed")
			break
		}
		fmt.Println("Channel is Open: ", ok, res)
	}
}

func initStrings(ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- "Go Guruji Manish"
	}
	close(ch)
}

// interview question
// buffered channels in golang async in nature, unbuffered are sync in nature
func Example7() {
	defer DisplayLine()

	fmt.Println("Example 7 buffered channels")

	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3
	// ch <- 4 // it will waited for extra sending
	close(ch)
	for d := range ch {
		fmt.Println("r: ", d)
	}
}

// more example of wait group in with channels
func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", i)

	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Worker %d done\n", i)
}

func Example8() {
	defer DisplayLine()
	fmt.Println("Example 8 simple example of wait group with channel")
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			worker(i, &wg)
		}(i)
		// go worker(i, &wg)
	}
	wg.Wait()
}

// interview question
// same above function using channel synchronization to wait go routines
func worker2(i int, check chan bool) {
	fmt.Printf("Worker2 %d starting\n", i)

	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Worker2 %d done\n", i)
	check <- true
}

func Example9() {
	defer DisplayLine()
	fmt.Println("Example 9 waiting with channel synchronization")

	chan1 := make(chan bool)
	for i := 1; i <= 5; i++ {
		// go func(chan bool) {
		// 	worker2(i, chan1)
		// }(chan1)
		go worker2(i, chan1)
		<-chan1
	}
}

// use of sync.Once in without channle
func shopping(p int) int {
	return p - 10
}
func saving(p int) {
	fmt.Println("Saving: ", p)
}
func Example10() {
	defer DisplayLine()
	fmt.Println("Example 10 with sync.Once")
	paise := 0
	var once sync.Once
	for i := 0; i < 3; i++ {
		paise += 10

		once.Do(func() {
			paise = shopping(paise)
		})
		saving(paise)
	}
}

// example to explain select statements with receiving value from multiple channels channels
func functionOne(chanl1 chan string) {
	time.Sleep(time.Second)
	chanl1 <- "Welcome to the creative development world of the Golang"
}
func functionSecond(chanl2 chan int) {
	time.Sleep(time.Second)
	chanl2 <- 220
}
func Example11() {
	defer DisplayLine()
	fmt.Println("Example 11 channels with select statements receiving values from multiple channels")
	chanl1 := make(chan string)
	chanl2 := make(chan int)

	go functionSecond(chanl2)
	go functionOne(chanl1)

	select {
	case val1 := <-chanl1:
		fmt.Println(val1)
	case val2 := <-chanl2:
		fmt.Println(val2)
	}
}

// basic example
func multiplyWithChannel(ch chan int) {
	fmt.Println(100 * <-ch)
}

func Example12() {
	defer DisplayLine()
	fmt.Println("Example 12 basic channel")

	ch := make(chan int)
	fmt.Println("Hello from example 12")
	go multiplyWithChannel(ch)
	ch <- 10
	fmt.Println("Bye from example 12")
}
