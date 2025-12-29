package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go fileProducer(ch)
	ProcessFile(ch)
}
func fileProducer(ch chan string) {
	files := []string{
		"a.txt",
		"b.txt",
		"c.py",
		"d.java",
		"h.go",
	}
	for _, fi := range files {
		fmt.Println("Process", fi)
		ch <- fi
	}

	// Channel sending signal completion automatically
	close(ch) // Only close(ch) communicates “no more values will ever come”
	// The above line says, no more data left to send through the channel and main goroutine can exit the program easily
}

func ProcessFile(ch chan string) {
	for fi := range ch {
		fmt.Println("Processed", fi)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("All files processed..!!")
}
