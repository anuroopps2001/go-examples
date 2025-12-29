package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Job    int
	Square int
}

func main() {
	// Created variable to work with multiple goroutines using waitGroup package
	var wg sync.WaitGroup

	// Created 2 channels
	jobs := make(chan int)
	results := make(chan Result)

	// number of goroutines managed by waitgroup
	workers := 3

	// waitgroup starts 3 child goroutines and manage the lifecycle
	wg.Add(workers)

	// starting 3 goroutines for a specific function
	for i := 0; i < workers; i++ {
		go worker(i, jobs, results, &wg)
	}

	// Wait till all the child goroutines finish their job
	go func() {
		wg.Wait()
		close(results)
	}()

	// sending load through the 'jobs' channel
	go func() {
		for j := 0; j <= 10; j++ {
			jobs <- j
		}
		close(jobs) // sender closing the channle to over avoide receiver in waiting state which leads to "dead lock"
	}()

	// accesing the 'results' channel data
	for res := range results {
		fmt.Println("result", res.Job, "->", res.Square)
	}

	fmt.Println("all jobs processed and results collected")

}

func worker(id int, jobs chan int, results chan Result, wg *sync.WaitGroup) {
	// executes at the end function execution and reduce the count of child goroutine for successful execution
	defer wg.Done()

	//
	for job := range jobs { // range "channel" does NOT look ahead.
		// It receives values one-by-one as they arrive.
		fmt.Println("worker", id, "processing", job)
		results <- Result{ // appending values into struct
			Job:    job,       // first iteration Job: 0
			Square: job * job, // Square: 0 * 0 = 0
		}
	}

	fmt.Println("worker", id, "exiting")
}
