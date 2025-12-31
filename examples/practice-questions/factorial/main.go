package main

import (
	"fmt"
	"log"
)

func main() {
	var input int

	fmt.Print("Enter a number to get its factorial: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatal("Please Enter an number to get It;s Factorial value..!!")
	}

	res, err := calFact(input)
	if err != nil {
		log.Fatalf("Error for an input %v: %v", input, err)
	}

	fmt.Printf("The factorial of %d is %d", input, res)
}

func calFact(number int) (uint64, error) {
	var fact uint64 = 1

	if number < 0 {
		return 0, fmt.Errorf("Negative numbers are not allowed")
	} else if number == 0 {
		return 1, nil
	}

	for num := 1; num <= number; num++ {
		fact = fact * uint64(num)
	}
	return fact, nil
}
