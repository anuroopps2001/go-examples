package main

import "fmt"

func countDown(n int) {
	if n <= 0 {
		fmt.Println("Invalid Number")
	} else {
		for i := n; i >= 1; i-- {
			fmt.Println(i)
		}
	}
}

func main() {
	countDown(7)
}
