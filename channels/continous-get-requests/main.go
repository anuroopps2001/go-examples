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

		// Function literal := Define a function here and run it in a goroutine immediately
		go func(link string) { // child goroutine can work on link passed by main goroutin
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // Passing 'l' for loop as argument into funtion literal i.e anonymus function
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
