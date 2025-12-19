package main

import "fmt"

type IntSlice []int

func (i IntSlice) sum() {
	value := 0
	for _, val := range i {
		value = value + val
	}
	fmt.Println(value)
}
