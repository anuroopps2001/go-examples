package main

import (
	"fmt"
	"time"
)

func main() {
	// channel created to allow struct type data communication through the channel
	stopCh := make(chan struct{})
	go worker(stopCh)

	time.Sleep(5 * time.Second)
	fmt.Println("Main sending stop signal..")
	close(stopCh)

	time.Sleep(1 * time.Second)
	fmt.Println("Program exited")

}

func worker(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Worker stopping..")
			return
		default:
			fmt.Println("worker doing work..")
			time.Sleep(1 * time.Second)
		}
	}
}
