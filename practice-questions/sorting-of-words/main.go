package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter comma separated words: ")
	input, _ := reader.ReadString('\n')
	inputformat := strings.TrimSpace(input)
	parts := strings.Split(inputformat, ",")
	sortedLetters(parts)
}

func sortedLetters(s []string) {
	slices.Sort(s)
	fmt.Print("Sorted in alphabetic order: ", s)
}
