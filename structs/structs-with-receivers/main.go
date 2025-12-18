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

	// firstPerson Variable holding the value [ REMINDER FOR POINTERS IN "GO"]
	firstPerson := person{
		firstName: "Steve",
		lastName:  "Smith",
		contactInfo: contactInfo{ // contactInfo contactInfo struct definition
			email:   "smudge@aussie.com",
			zipCode: 577540,
		},
	}

	/* // Below &firstPerson statement is saying that, give me memory address of firstPerson variable which is storing the struct
	// and firstPersonPointer will point to memory address of firstPerson variable and not to the actual struct firstPerson variable storing
	// Remember that, firstPerson variable actually gives value stored, but by using &firstPerson, we saying, "give me the memory address and not the actual value"
	// firstPersonPointer actual holds memory address give by firstPerson and not the actual value and this can be called as pointer
	// firstPersonPointer holds address and firstPerson holds the value
	// firstPerson holds value and firstPersonPointer holds memory
	// To turn, value into memory address we do, &value and exactly below does */
	firstPersonPointer := &firstPerson

	// Go does this behind the scenes: (&firstPerson).updateFirstName("Jamie")
	// firsPerson.updateFirstName("Jamie") also does the same thing
	firstPersonPointer.updateFirstName("Jamie") // calling function with pointerVariable which holds the memory address

	// Now we can the receiverTypeFunction with an variable which holdes the values of custom type 'person'
	firstPerson.receiverTypeFunction()
}

// This function will update the firstName of an person
// with '*' here means, inside this method, I want to modify the ORIGINAL struct.
// here, receiver "*person" will be called as type "POINTER TO PERSON"
func (pointerToPerson *person) updateFirstName(newFirstName string) {

	// turn memory address into value, *address
	// *<memory_address OR pointer>, this is like saying, take this memory address and give the value stored at that specific memory address
	(*pointerToPerson).firstName = newFirstName // updating the firstName of person custom type with the value passed while calling this function

	// pointerToPerson.firstName = newFirstName this also works, Go automatically does (*pointerToPerson).firstName = newFirstName
}

// functions with receivers of type 'person' which is an custom type
func (p person) receiverTypeFunction() {
	fmt.Printf("%+v", p)
}
