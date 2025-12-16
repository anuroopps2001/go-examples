package main

import "fmt"

func sumUpTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func main() {
	Sum := sumUpTo(5)

	fmt.Println(Sum)
}
