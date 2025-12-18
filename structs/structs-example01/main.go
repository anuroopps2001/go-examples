package main

import "fmt"

// Create a custom type named 'person' of which underlying data type is 'struct'
type person struct {
	firstName string
	lastName  string
}

type Author struct {
	Name   string
	Branch string
	Year   int
}

type HR struct {
	Human   Author // Embedded structs, here Human is of custom type HR
	Country string
}

func main() {
	// assigning values into struct of custom type person := FIRST WAY OF ASSIGINING VALUES INTO 'struct' i.e DECLARING STRUCT
	person01 := person{firstName: "Alex", lastName: "Anderson"}

	/* // assigning values into struct of custom type person := SECOND WAY OF ASSIGNING VALUES INTO 'struct' i.e DECLARING STRUCT
	// Because, now the order will be Jamie will be assigned to firstName and Smith to lastName and if I change order inside the type definition
	// then lastName will be Jamie and firstName will be Smith */
	person02 := person{"Jamie", "Smith"} // when we assigin this way, we cannot order values variables inside 'person' custom type definition

	// // assigning values into struct of custom type person := THIRD WAY OF ASSIGNING VALUES INTO 'struct' i.e DECLARING STRUCT
	var person3 person // Go will create a variable named person3 of type 'person' and assigns something called "ZERO VALUES"
	// when variables are defined of type struct and not assigned with any values.

	// Assigning values into declated variables of struct type
	person3.firstName = "Ben"
	person3.lastName = "Stokes"
	fmt.Println(person3)       // basically retunrs { } which has 2 empty strings
	fmt.Printf("%+v", person3) // returns {firstName: lastName:} where 2 fields contain empty strings

	/*
	   for string type of values, zero-value = ""
	   for int type of values, zero-value = 0
	   for float type of values, zero-value = 0
	   for bool type of values, zero-value = false
	*/
	// It will print {Alex Anderson} {Jamie Smith} similar to dict type in PYTHON
	fmt.Println(person01, person02)

	// Values assignment into embedded structs
	result := HR{
		Country: "USA",
		Human: Author{
			Name:   "Dhruv",
			Branch: "IT",
			Year:   2024,
		},
	}

	fmt.Printf("Details of Embedded struct Human HR: %+v", result)
}
