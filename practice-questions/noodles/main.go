package main

import (
	"fmt"
)

func main() {

	words := []string{"hello", "world", "leetcode"} //slice of strings
	sortWords(words)                                // passing slice of strings while calling function
	fmt.Println(words)

}

func countDistincVowels(word string) int {
	vowels := map[byte]bool{
		// assigining values into maps
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	// make a map with key being byte type and value being bool type
	seen := make(map[byte]bool)

	for i := 0; i < len(word); i++ {
		ch := word[i]
		if vowels[ch] {
			seen[ch] = true
		}
	}

	// Return length of unique vowels from an give string
	return len(seen)
}

func sortWords(words []string) { // Now words := []string{"hello", "world", "leetcode"}
	n := len(words) // n = 3

	for i := range n { // can be written as for i := 0; i < n ; i++
		for j := i + 1; j < n; j++ { // can be written as for j := range n
			countI := countDistincVowels(words[i]) // get the count of distinct vowels, by passing an string from slice of strings
			countJ := countDistincVowels(words[j]) // get the count of distinct vowels, by passing an string from slice of strings

			if countI > countJ || (countI == countJ && words[i] > words[j]) { // words[i] > words[j] == Checks Does words[i] come after words[j] alphabetically?
				words[i], words[j] = words[j], words[i]
			}
		}
	}
}
