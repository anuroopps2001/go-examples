package main

import "fmt"

func main() {
	arrayElements := IntSlice{4, 1, 2, 1, 2} // Because IntSlice = []int{} i.e slice of integers

	fmt.Println(arrayElements.uniqueNumbers())
}
