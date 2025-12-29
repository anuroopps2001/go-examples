package main

import "fmt"

func main() {
	ch := make(chan int)
	/* When multiple goroutines send to the same channel,
	   you cannot guarantee which value will be received first */
	go sliceSum([]int{1, 2, 3}, ch)
	go sliceSum([]int{4, 5, 6}, ch)
	fmt.Println("Received total: ", <-ch)
	fmt.Println("Received total: ", <-ch)

}

func sliceSum(s []int, c chan int) {
	total := 0

	for _, num := range s {
		total += num
	}
	fmt.Println("sending: ", total)
	c <- total
}
