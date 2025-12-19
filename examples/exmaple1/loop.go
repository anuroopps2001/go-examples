package main

func getUnique(u []int) []int {
	finalOut := []int{}
	for i := 0; i < len(u); i++ {
		count := 0
		for j := 0; j < len(u); j++ {
			if u[i] == u[j] {
				count++
			}
		}
		if count == 1 {
			finalOut = append(finalOut, u[i])
		}
	}
	return finalOut
}
