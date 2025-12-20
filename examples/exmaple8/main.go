package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{name: "Anuroop", age: 26}
	p.birthDay()
	fmt.Println("After calling birthDay():", p.age) // age should stay 26

	p.haveBirthDay()
	fmt.Println("After calling haveBirthDay():", p.age) // age should become 27

}

// Value receiver → works on a copy
func (pe Person) birthDay() {
	pe.age = pe.age + 1
	fmt.Println("age is still", pe.age)
}

// Pointer receiver → modifies original
func (pe *Person) haveBirthDay() {
	pe.age = pe.age + 1
	fmt.Println("age becomes", pe.age)
}
