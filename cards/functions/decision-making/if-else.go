package main

import "fmt"

func checkAge(age int) {
	if age >= 18 && age < 21 {
		fmt.Println("Being a man, cannot marry yet..!!!")
	} else if age >= 21 && age <= 40 {
		fmt.Println("Being a man you allowed to marry")
	} else {
		fmt.Println("Strictly not allowed..!!")
	}
}
