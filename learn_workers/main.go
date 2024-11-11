package main

import (
	"fmt"
	"math/rand"
	"time"
)

// interview question

// example to example workers using go routines and channel
// 3 workers 100 tasks with struct and print completed tasks

type Job struct {
	id        int
	completed bool
}

func worker(id int, jobs <-chan Job, result chan<- Job) {
	for j := range jobs {
		fmt.Println("worker: ", id, "started job: ", j.id)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("worker: ", id, "finished job: ", j.id)

		j.completed = true
		result <- j
	}
}

func main() {
	start := time.Now()
	fmt.Println("Starts...")
	tasks := 100
	workers := 10 //5

	jobs := make(chan Job, tasks) // buffered channel for tasks
	results := make(chan Job, tasks)

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
		if res.completed {
			completed++
		}
	}
	fmt.Println("Completed Tasks: ", completed)
	fmt.Println("Ends...")
	end := time.Since(start)
	fmt.Println("Time Taken: ", end)

}
