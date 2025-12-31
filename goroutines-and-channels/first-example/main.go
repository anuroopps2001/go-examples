package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amamzon.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://udemy.com",
	}

	// channel creation to allow communicatiin with type of data being 'string'
	// string type channels allow, communication between goroutines is expected to be of type string
	c := make(chan string)

	for _, link := range links { // start multiple child go routines to work on different links by calling checkLink function
		// Pass the channel variable from where main goroutine is creating the child goroutines
		go checkLink(link, c) // go keyword is used to create child goroutines by main goroutine
	}

	for i := 0; i < len(links); i++ {
		// Receiving message from channel onto to the terminal
		fmt.Println(<-c) // Blocking line of code in general i,e main goroutine will wait for the messages through the channel
	}

	// fmt.Println(<-c) // This will never complete, because there is no more data coming through the channels, so main goroutine never completes

}

func checkLink(s string, c chan string) {
	_, err := http.Get(s)

	if err != nil {
		fmt.Println("Error", err)
		// sending data into an channel 'c' and the syntax is :=  channel_name <- <data_we_want_to_send_into_channel>
		c <- "Might be down"
		return
	}
	fmt.Println(s, "connected..")
	// sending message into channel 'c'
	c <- "Yup..It's working"
}
