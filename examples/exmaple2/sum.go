package main

type IntSlice []int

func (i IntSlice) sum() int {
	value := 0
	for _, val := range i {
		value = value + val
	}
	return value
}
