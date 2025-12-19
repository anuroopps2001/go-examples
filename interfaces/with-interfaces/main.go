package main

import "fmt"

// Any type that has a method getGreeting() string is considered a bot.
type bot interface {
	getGreeting() string // Methods having name getGreeting() and with return type being 'string' are managed by this specific interface
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	// spanishBot receiver is used with Method getGreeting() string → OK, it satisfies bot.
	sb := spanishBot{}
	// englishBot receiver is used with Method getGreeting() string → OK, it satisfies bot.
	eb := englishBot{} // eb is the variable, which can call the methods accepting the values of type englishBot

	fmt.Println(eb.getGreeting())

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) { // You must pass a value whose type satisfies the bot interface
	fmt.Println(b.getGreeting())
}

// This without short name like 'eb' works fine
func (englishBot) getGreeting() string { // This specific method can be called by varibales holding the values of type englishBot

	return "Hi there..!!"
}

// This also works as we seen in many examples
func (sb spanishBot) getGreeting() string {
	return "Hola..!!"
}
