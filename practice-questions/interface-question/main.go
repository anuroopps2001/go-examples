package main

import "fmt"

func main() {
	wantedString := "HELLO WORLD!"

	ReadString()
	gotString := PrintString()

	if gotString != wantedString {
		fmt.Println(fmt.Errorf("Error: ReadString() = %v, but wanted %v", gotString, wantedString))
	}

	fmt.Println(gotString)
}
