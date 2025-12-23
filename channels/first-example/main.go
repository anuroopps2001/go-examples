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

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(s string) {
	_, err := http.Get(s)

	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(s, "connected..")
}
