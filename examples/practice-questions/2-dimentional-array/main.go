package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter 2 numbers separated by comma:")
	input, _ := reader.ReadString('\n')
	inputnumbers := strings.TrimSpace(input)

	parts := strings.Split(inputnumbers, ",") // slice of strings

	X, _ := strconv.Atoi(parts[0])
	Y, _ := strconv.Atoi(parts[1])

	output := twoDimArray(X, Y)
	fmt.Println(output)

}

func twoDimArray(x, y int) [][]int {
	output := make([][]int, x) // Make a 2-dim slice of integers of size 'x'

	for num1 := range x {
		output[num1] = make([]int, y)
		for num2 := range y {
			output[num1][num2] = num1 * num2 // Ex:- [[0,0,0,0,0],[0,1,2,3,4],[0,2,4,6,8]]
		}
	}
	return output
}
