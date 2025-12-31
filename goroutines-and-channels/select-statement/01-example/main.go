package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(6 * time.Second)
		ch1 <- "from channel1 after 1 second"
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- "from channel2 after 5 seconds" // becomes ready after 5 seconds and sends data first
	}()

	select { // select blocks until at least one case is ready.

	/* If both are ready:
	    Go chooses one randomly
		This prevents starvation
		You must not depend on order */
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2: // since data comes first through this specific channel, this case will be executed and other
		// case statement will be ignored
		fmt.Println(msg2)

		/* Wait for message
		   OR wait for 1 seconds
		   Whichever happens first wins */
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("I won't wait until channels send data, so I will execute myself immediately")
		// select without default → waits
		// select with default → never waits
	}
	fmt.Println("Received data and exiting the program")
}
