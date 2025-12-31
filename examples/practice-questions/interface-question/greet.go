package main

import (
	"bufio"
	"os"
	"strings"
)

var input string

func ReadString() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}
}

func PrintString() string {
	return strings.ToUpper(input)
}
