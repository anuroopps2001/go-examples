package main

import "fmt"

func primeNumbers(numbers int) {
	for i := 1; i <= numbers; i++ {
		fmt.Println(i)
	}
}

func main() {
	primeNumbers(5)
}
