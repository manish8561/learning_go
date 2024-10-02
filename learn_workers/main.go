package main

import (
	"fmt"
	"math/rand"
	"time"
)

// example to example workers using go routines and channel
// 3 workers 100 tasks with struct and print completed tasks

type Job struct {
	id int
}

func worker(id int, jobs <-chan Job, result chan<- bool) {
	for j := range jobs {
		fmt.Println("worker: ", id, "started job: ", j.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("worker: ", id, "finished job: ", j.id)

		result <- true
	}
}

func main() {
	fmt.Println("Starts...")
	tasks := 100
	workers := 5

	jobs := make(chan Job, tasks)
	results := make(chan bool)

	for i := 1; i <= workers; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= tasks; j++ {
		jobs <- Job{id: j}
	}
	close(jobs)

	completed := 0
	for r := 1; r <= tasks; r++ {
		res := <-results
		if res {
			completed++
		}
	}
	fmt.Println("Completed Tasks: ", completed)
	fmt.Println("Ends...")
}
