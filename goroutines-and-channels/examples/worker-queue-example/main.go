package main

import (
	"fmt"
	"time"
)

func main() {

	firsCh := make(chan int)
	secondCh := make(chan struct{})
	go worker(1, firsCh, secondCh) // new child goroutine is started
	go worker(2, firsCh, secondCh) // another child goroutine started

	go func() { // It;s an "Funtion literal", one more child goroutine started and main goroutine is free now and executes it's work
		for i := 0; i < 1700; i++ {
			firsCh <- i
		}
	}()
	time.Sleep(6 * time.Second)          // Main goroutine will sleep for 6 seconds
	fmt.Println("Main stopping workers") // After 6 seconds starts it;s execution
	close(secondCh)                      // sending graceful shutdown signal to all the child goroutines after 6 seconds of sleep
	fmt.Println("Program exited..!")

}

func worker(id int, fi chan int, se chan struct{}) {
	for {
		select {
		case <-se:
			fmt.Println("Worker", id, "Stopped..")
			return
		case item := <-fi:
			fmt.Println("Worker", id, "Processing Item: ", item)
			time.Sleep(1 * time.Second)
		}
	}
}
