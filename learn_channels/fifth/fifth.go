package fifth

import (
	"fmt"
	"demo/examples"
	// "time"
)

// interview question for sending and receiving value using sending and receiving channel and channel for synchronization
func sending(ch chan<- int){
	fmt.Println("Sending value")
	for i:=1; i<=5; i++{
		ch <-i
	}
	close(ch)
}

func receiving(ch <-chan int, check chan bool){
	for v := range ch{
		fmt.Println("Receiving the value: ", v)
	}
	check <- true
}

func Example(){
	defer examples.DisplayLine()

	fmt.Println("Fifth Example for 2 goroutine with one channel")

	ch := make(chan int)
	check := make(chan bool)
	go sending(ch)
	go receiving(ch,check)
	// time.Sleep(1 * time.Second)
	<-check
}
