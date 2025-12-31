package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var numbers string
	fmt.Println("Enter a sequence of comma-separated numbers: ")
	fmt.Scan(&numbers)

	filledSlice := intputSlice(numbers)

	fmt.Println(filledSlice)

}

func intputSlice(n string) []int {
	// Define empty slice of integers to hold comma separated values
	emptySlice := []int{}

	// Create a slice of strins using Split method with each string being comma separated
	parts := strings.Split(n, ",")

	// Convert string of an slice of strings into number
	for _, num := range parts {
		n, err := strconv.Atoi(num)

		if err != nil {
			continue
		}

		emptySlice = append(emptySlice, n) // Append into slice of integers
	}
	return emptySlice
}
