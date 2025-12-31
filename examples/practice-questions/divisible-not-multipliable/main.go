package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	firstRange := 2000
	secondRange := 3200

	numbersOut := divisibleNotMultiple(firstRange, secondRange)
	fmt.Print(strings.Join(numbersOut, ","))
}

func divisibleNotMultiple(firstRange int, secondRange int) []string {

	var numbers []string
	for firstRangeval := firstRange; firstRangeval <= secondRange; firstRangeval++ {
		// check whether number is divisible by 7 and not by 5
		if firstRangeval%7 == 0 && firstRangeval%5 != 0 {
			numbers = append(numbers, strconv.Itoa((firstRangeval)))
		}
	}
	return numbers
}
