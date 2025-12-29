package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go produceLog(ch)
	receiveLog(ch)

}

func produceLog(ch chan string) {
	logs := []string{
		"auth success",
		"db connected",
		"cache warmed",
	}
	for _, st := range logs {
		ch <- st
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func receiveLog(ch chan string) {
	for log := range ch {
		fmt.Println("Processing", log, "log")
	}
	fmt.Println("No more logs to process")
}
