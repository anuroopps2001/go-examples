package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make(chan int, 5) // buffered channel with capacity of 10
	stopCh := make(chan struct{})
	var wg sync.WaitGroup // WaitGrop is used to wait for a group of goroutines or tasks to finish.

	for i := 1; i <= 2; i++ {
		wg.Add(1) // Add one new goroutine for every iteration
		go worker(i, queue, stopCh, &wg)
	}

	// Function literal
	go func() {
		for i := 0; i < 100; i++ {
			queue <- i
		}
	}()

	time.Sleep(5 * time.Second)
	close(stopCh)
	wg.Wait()
	fmt.Println("All workers stopped. Program exiting.")
}
