package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, queue chan int, stopCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stopCh:
			fmt.Println("Worker", id, "stopped")
			return

		case item := <-queue:
			fmt.Println("Worker", id, "processing", item)
			time.Sleep(1 * time.Second)
		}
	}
}
