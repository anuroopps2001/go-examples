package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	checkAge(age)
}
