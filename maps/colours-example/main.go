package main

import (
	"fmt"
)

func main() {

	// Declaring a map and assigning a values := FIRST WAY
	colours := map[string]string{ // syntax is map[key_data_type]value_data_type
		// So, in our exmaple, key will be of string data_type and value will also be of string data_type
		"red":   "#ff0000", // key=red which is of string type and value=#ff000 is also an string type
		"green": "#008000",
	}
	fmt.Println("Before deleting an key-value pair; ", colours)
	delete(colours, "red") // Deleting an key-value from an map
	fmt.Println("After deleting an key-value pair; ", colours)
	colours["pink"] = "#265ed" // Assigning key-value pair into existing map

	// Declaring a map := SECOND WAY
	var names map[string]string // Just declaring a map

	names = make(map[string]string) // Initializing a earlier declared map
	names["first_name"] = "Anuroop" // Assigining
	names["last_name"] = "P S"
	fmt.Println(names)

	// Declaring a map := THIRD WAY
	places := make(map[string]string) // Declaration and initialization
	places["Bengaluru"] = "Karnataka" // Assignment
	places["Delhi"] = "NCR"

	// Deleting key-value pairs from a map using delete builtin function
	fmt.Println("Before deleting an key-value pair; ", places)
	delete(places, "Delhi") // syntax is delete(map_name, <key_from_where_key-value_pair_to_be_deleted)
	fmt.Println("After deleting an key-value pair; ", places)

	printMap(colours)
}

func printMap(col map[string]string) {
	for color, hex := range col {
		fmt.Println("Hex code of color", color, "is", hex)
	}
}
