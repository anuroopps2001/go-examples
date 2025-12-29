package main

import "fmt"

func main() {

	// Creating an channel with var name 'ch'
	ch := make(chan string)

	// goroutine and function literal
	go func() {
		ch <- "Done"
	}()

	msg := <-ch // msg receives message through the channel
	fmt.Println(msg)
}
