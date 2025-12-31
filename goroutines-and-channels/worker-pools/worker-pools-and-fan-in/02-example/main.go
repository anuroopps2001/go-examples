package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()

	defer func() {
		duration := time.Since(startTime)
		fmt.Printf("Time taken for Program execution: %s\n", duration)
	}()

	results := make(chan int)

	var wg sync.WaitGroup

	producers := 3
	wg.Add(producers) // look for 3 child goroutines

	// start 3 new child goroutines
	for i := 1; i <= producers; i++ {
		go worker(i, results, &wg)
	}

	// wait for all senders to finish data from 3 child goroutines and close the channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// single goroutine is receiving the data sent by multiple sender goroutines into single channel and this is called as
	// fan-in function
	for val := range results {
		fmt.Println("received", val)
	}

	fmt.Println("all processes are done and channel closed")
}

func worker(goroutineId int, response chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(goroutineId) * time.Second) // blocks only the current goroutine, not the whole program.

	response <- goroutineId * 10 //3 child goroutines will send data into same channel
	fmt.Println("goroutine", goroutineId, "sent", goroutineId*10)
}
