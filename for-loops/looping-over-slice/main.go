package main

import "fmt"

func main() {
	usrNames := []string{"anuroop", "satisg", "ramesh"}

	for _, name := range usrNames { // if index is not required, then use underscore (_)
		fmt.Println(name)
	}
}
