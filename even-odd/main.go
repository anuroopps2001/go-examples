package main

import "fmt"

func main() {
	integers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, integer := range integers {
		if integer%2 == 0 {
			fmt.Println(integer, "is an even number")
		} else {
			fmt.Println(integer, "is an odd number")
		}
	}
}
