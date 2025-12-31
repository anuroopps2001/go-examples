package main

import "fmt"

func sunOfEven(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum = sum + i
		}
	}
	return sum
}

func main() {
	evensum := sunOfEven(6)

	fmt.Println(evensum)
}
