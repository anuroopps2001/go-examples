package main

import "fmt"

func printEvenOdd(n int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {
			fmt.Println(i, "Odd")
		}
	}
}

func main() {
	printEvenOdd(3)
}
