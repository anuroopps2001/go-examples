package main

import (
	"math/rand"
)

func generatePassword(length int) string {
	// Allowed characters for the password

	allwoedCharacters := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_"

	// Initializing password with empty string
	password := ""

	for i := 0; i < length; i++ {

		// Generate a random index
		randomIndex := rand.Intn(len(allwoedCharacters))

		// Get the character from string using the above generated index
		char := allwoedCharacters[randomIndex]

		// Append into the defined empty string
		password = password + string(char)
	}

	return password
}
