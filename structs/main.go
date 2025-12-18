package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // equivalant to contact contactInfo OR contactInfo contactInfo and it's an EMBEDDED STRUCT
}

func main() {
	firstPerson := person{
		firstName: "Steve",
		lastName:  "Smith",
		contactInfo: contactInfo{ // contactInfo contactInfo struct definition
			email:   "smudge@aussie.com",
			zipCode: 577540,
		},
	}

	fmt.Printf("%+v", firstPerson)
}
