package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	data := make(chan int)

	go producer(data, ctx)

	go worker(data, ctx)

	time.Sleep(200 * time.Microsecond)

	fmt.Println("main: cancelling  context")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("main: exited cleanly")
}

func producer(data chan int, ctx context.Context) {
	for i := 0; i <= 10; i++ {
		select {
		case data <- i:
			fmt.Println("producer sent:", i)
			time.Sleep(300 * time.Microsecond)
		case <-ctx.Done():
			fmt.Println("producer: received cancel, signal")
			return
		}
	}
}

func worker(data chan int, ctx context.Context) {
	for {
		select {
		case v, ok := <-data:
			if !ok {
				fmt.Println("worker: data channel closed, exiting")
				return
			}
			fmt.Println("worker processing", v)
			time.Sleep(500 * time.Microsecond)

		case <-ctx.Done():
			fmt.Println("worker: received stop signal, exiting")
			return
		}
	}
}
