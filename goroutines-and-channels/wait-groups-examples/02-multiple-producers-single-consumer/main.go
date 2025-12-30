package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup

	producers := 3
	wg.Add(producers)

	for i := 1; i <= producers; i++ {
		go producer(i, ch, &wg)
	}
	// Function literal
	go func() { // waiter child goroutine
		wg.Wait()
		close(ch) // since multiple goroutines sending data, we need to close the channel on the new goroutine,
		// where wait() is functional
	}()

	for v := range ch { // Receiving data from channels
		fmt.Println("consumed", v)
	}

	fmt.Println("all producers are done, channel closed")
}

func producer(id int, c chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 3; i++ {
		value := id*10 + 1
		fmt.Println("producer", id, "sending", value)
		c <- value
	}
}
