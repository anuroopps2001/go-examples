package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {
	sb := spanishBot{}
	eb := englishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

// Because we are using the same function i.e printGreeting() twice, even though we are using different args to call those funcitons, we will get an error
func printGreeting(eb englishBot) { // says, this function needs to be called with the function of type englishBot which is basically of struct type
	// Remember custom type we defined earlier and 'eb' is arguement which we can use inside the function
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi there..!!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola..!!"
}
