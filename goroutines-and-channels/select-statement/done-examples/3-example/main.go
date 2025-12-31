package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	data := make(chan int)

	go producer(data, done)
	go worker(data, done)
	fmt.Println("Main initializing shutdown")
	time.Sleep(3000 * time.Microsecond)
	close(done) // Drop everything and exit immediately.
	time.Sleep(1 * time.Second)
	fmt.Println("main exited cleanly")

}

func producer(data chan int, done chan struct{}) {
	for i := 0; i <= 10; i++ {
		select {
		case data <- i:
			fmt.Println("producer sent:", i)
			time.Sleep(300 * time.Microsecond)
		case <-done:
			fmt.Println("producer: shutdown")
			return
		}
	}
}

func worker(data chan int, done chan struct{}) {
	for {
		select {
		case v := <-data:
			fmt.Println("worker processing:", v)
			time.Sleep(500 * time.Microsecond)

		case <-done:
			fmt.Println("worker: shutdown")
			return
		}
	}
}
