package main

import "fmt"

func printTable(n int) {
	// product := 0
	for i := 1; i <= 10; i++ {
		fmt.Println(n, "*", i, "=", i*n)
	}
	// fmt.Println(product)
}

func main() {
	printTable(5)
}
