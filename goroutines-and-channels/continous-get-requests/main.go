package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amamzon.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://udemy.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c) // Blocking line of code
	}

	/* 	for { // Infinite loop
		go checkLink(<-c, c) // creating new goroutine
	} */

	for l := range c { // Here, main goroutine can work on links one by one

		// FUNCTION LITERALS
		// We are using function literals here to not to block or pause the main routine
		// Function literal := Define a function here and run it in a goroutine immediately, when channel receives data without blocking the main goroutine
		go func(li string) { // child goroutine can work on link passed by main goroutin and that is passed as argument into function literal
			time.Sleep(5 * time.Second) // integrating a pause between creating an each goroutine
			checkLink(li, c)            // calling checkLink function again and again with new goroutines
		}(l) // l, is value held by main routine and since this function literal is running a separate child goroutine,
		// we need to make child goroutine understand this by the same same value as arguement i,e PASS BY VALUE
	}
}

func checkLink(s string, c chan string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println("Error:", err)
		c <- s // passing link again into the channel
		return
	}
	fmt.Println(s, "connected..!!")
	c <- s
}
