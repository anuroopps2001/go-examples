package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		fmt.Println("worker: waiting for done")
		<-done
		fmt.Println("worker: received done signal, exiting")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("main: closing done")
	close(done)

	time.Sleep(1 * time.Second)
}
