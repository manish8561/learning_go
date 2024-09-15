package main

import (
	"fmt"
	"sync"

	"exmaple.com/demo/d2/examples"
)

/* reading and writing files using channels */
func main() {
	var wg sync.WaitGroup
	dataCh := make(chan string)

	wg.Add(2)
	go examples.Reading(dataCh, &wg)
	go examples.Writing(dataCh, &wg)

	wg.Wait()

	fmt.Println("Printing from channels")

}
