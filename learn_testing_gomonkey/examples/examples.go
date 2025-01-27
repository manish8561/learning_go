package examples

import (
	"fmt"
	"time"
)

func wait() {
	fmt.Println("Waiting ..............")
	time.Sleep(5 * time.Second)
}

func Add(x, y int) int {
	wait()
	return x + y
}
