package main

import (
	"fmt"
	"time"
)

func main() {
	go hello() // This line will start child goroutine

	time.Sleep(10 * time.Second) // We use 1 second only to keep main alive long enough for the goroutine to run;
	//  the exact duration does not matter and is only for learning purposes.

}

func hello() {
	fmt.Println("Hello from goroutine")
}
