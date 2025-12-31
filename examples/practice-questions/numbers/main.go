package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	// A buffer reads a chunk of data once and keeps it in memory
	reader := bufio.NewReader(os.Stdin) // use bufio.NewReader when you want to read a full line (especially comma-separated input).
	fmt.Print("Enter comma-separated numbers: ")

	// Read input until you see a newline character (\n)
	input, _ := reader.ReadString('\n') // input == "1,2,3,4\n"

	/* strings.Trimspace removes:
	leading spaces
	trailing spaces
	newlines (\n)
	tabs */
	input = strings.TrimSpace(input)

	// "1,2,3,4" → []string{"1", "2", "3", "4"}
	parts := strings.Split(input, ",") // Create a slice of strings separated by commas
	output := calc(parts)
	fmt.Println(output)

}

func calc(d []string) []int {
	const c = 50
	const h = 30
	output := []int{}
	for i := 0; i < len(d); i++ {

		// converting each element of slice of strings to int, before using into required math expression
		num, _ := strconv.Atoi(d[i])

		// Formula used: Q = Square root of [(2 * C * D)/H]
		q := float64(((2 * c * num) / h))
		sqrt := math.Sqrt(q)
		roundedsqrt := math.Round(sqrt)
		output = append(output, int(roundedsqrt))

	}
	return output
}
