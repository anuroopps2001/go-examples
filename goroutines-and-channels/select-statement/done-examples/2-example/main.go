package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}) // carries no values, only a signal

	data := make(chan int) // carries actual values

	go func() {
		for {
			select {
			case v := <-data:
				fmt.Println("worker received", v)
			case <-done:
				fmt.Println("worker received done signal, exiting")
				return
			}
		}
	}()

	data <- 1
	time.Sleep(1 * time.Second)

	data <- 2
	time.Sleep(2 * time.Second)

	fmt.Println("main: closing done")
	close(done)

	data <- 3 // Once you signal done, you must stop sending work.
	time.Sleep(3 * time.Second)
	fmt.Println("main: exiting")
}
