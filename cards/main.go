/* package main

func main() {

	// var <variable_name> <data_type_to_this_variable> = assigned_value_to_variable
	// GO is an statically typed programming language, i.e we cannot change the variables data types randomly

	// var card string = "Ace of Spades"    // one way of defining variables

	// card := "Ace of spades" // second way of defining variables for the first time.

	// card = "Five of Diamonds" // while re-assigining the value to the already defined variable, this is the syntax

	// card := newCard()

	// cards := []string{"Ace of Diamonds", newCard()}

	// cards := deck{"Ace of Diamonds", newCard()}

	deckOfCards := newDeck()
	// cards = append(cards, "Six of Spades")

	/* for i, card := range cards {
		fmt.Println(i, card)
	} */
// cards.printCards()
/*
	hand, remainingCards := deal(deckOfCards, 5)

	hand.printCards()
	remainingCards.printCards()

} */

// func <func_name>s() <when executed this function, the value it return of data type "string">
/* func newCard() string {
	return "Five of Diamonds"
} */

package main

import "fmt"

func main() {
	// calling function by setting variable to function name, and because function not expecting any args, here no args are passed
	deckOfCards := newDeck()

	// deckOfCards is now slice
	fmt.Print(deckOfCards)

	// setting another variable for another function called countCards defined in deck.go
	totalNumberOfCards := countCards(deckOfCards)
	fmt.Println("Total number of cards: ", totalNumberOfCards)

	fmt.Println("The last card in the deck is: ", lastCard(deckOfCards))

	// storing the function output into variable by calling the function with required args
	firstCardOfDeck := firstCard(deckOfCards)

	fmt.Println("The first card in the deck is: ", firstCardOfDeck)

	printCards(deckOfCards)
}
