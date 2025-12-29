package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	job := make(chan int)
	workers := 3

	wg.Add(workers) // start 3 child goroutines

	for id := range workers {
		go worker(id, job, &wg) // create 3 goroutines and each calls worker function
	}

	// worker pool load
	for j := 0; j <= 10; j++ {
		job <- j
	}

	close(job)
	wg.Wait()
	fmt.Println("all jobs processed")
}

func worker(id int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // decrement one goroutine for each successful function call
	for c := range ch {
		fmt.Println("worker id", id, "processing", c)
	}
	fmt.Println("worker", id, "exiting")

}
