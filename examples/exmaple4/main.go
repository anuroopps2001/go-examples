package main

import "fmt"

func main() {
	var firstMap = make(map[string]int)
	firstMap["First"] = 1
	firstMap["Second"] = 2
	firstMap["Third"] = 3
	firstMap["Fourth"] = 4

	readKeyValue(firstMap)
}

func readKeyValue(values map[string]int) {
	keys := []string{}
	vals := []int{}
	for key, val := range values {
		keys = append(keys, key)
		vals = append(vals, val)
	}
	fmt.Println("Keys of this map are: ", keys)
	fmt.Println("Values of this map are: ", vals)
}
