package examples

import "fmt"

func ExamplFirst() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		ch1 <- 1
	}()
	go func() {
		ch2 <- 2
	}()
	select {
	case val := <-ch1:
		fmt.Println("Received", val, "from ch1")
	case val := <-ch2:
		fmt.Println("Received", val, "from ch2")
	}

}
