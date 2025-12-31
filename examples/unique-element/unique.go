package main

// Define a custom type to slice of integers
type IntSlice []int

func (u IntSlice) uniqueNumbers() IntSlice {

	// creating an slice of ints to hold uniqueElements
	uniqueElements := []int{}

	for _, val := range u { // taking the index at each loop
		count := 0
		for _, v := range u {
			if val == v { // comparing the value at each index of i and j
				count++
			}
		}

		if count == 1 {
			uniqueElements = append(uniqueElements, val)
		}
	}

	return uniqueElements // returning output of type IntSlice i.e []int{} slice of integers

}
