package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 1)
	n := 5
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println(v)
	}

}
