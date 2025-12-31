package main

import (
	"fmt"
	"log"
)

func main() {
	var n int

	fmt.Printf("Enter an number: ")
	_, err := fmt.Scan(&n)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resultMap := prodMap(n)
	fmt.Println(resultMap)

}
func prodMap(n int) map[int]int {
	outputMap := make(map[int]int)
	for i := 1; i <= n; i++ {
		// key = value
		outputMap[i] = i * i
	}

	return outputMap
}
