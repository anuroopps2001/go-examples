package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	p := Person{Name: "Anuroop"}
	p.Greet()
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is ", p.Name)
}
