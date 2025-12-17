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
	// totalNumberOfCards := countCards(deckOfCards)
	totalNumberOfCards := deckOfCards.countCards() // calling methods which are defined with receivers.
	fmt.Println("Total number of cards: ", totalNumberOfCards)

	// calling a method using receiver
	lastCardOfDeck := deckOfCards.lastCard()
	fmt.Println("The last card in the deck is: ", lastCardOfDeck)

	// storing the function output into variable by calling the function with required args
	//firstCardOfDeck := firstCard(deckOfCards)
	firstCardOfDeck := deckOfCards.firstCard()

	fmt.Println("The first card in the deck is: ", firstCardOfDeck)

	/* printCards(deckOfCards) */

	deckOfCards.printCards()

	// because along with dealOfDeck function, im not passing any receiver
	// dealOfDeck function returns 2 values when we call that and both are of type deck
	hand, remainingDeck := dealOfDeck(deckOfCards, 5)
	/* fmt.Println(hand)  // this is valid
	fmt.Println(remainingDeck)   // this is also valid */
	hand.printCards()
	remainingDeck.printCards()

	// converting string into slice of bytes i,e byte slice
	fmt.Println("Byte slice value of string is :", []byte("Hi there")) // byte slice also called as slice of bytes calculation

	fmt.Println(deckOfCards.toString())

	// to save a list of cards into local system
	deckOfCards.saveTofile("my_decK_of_cards")

	// calling newDeckFromFile function
	// calling newDeckFromFile("my_decK_of_cards") this way works, but it doesn't print the values properly
	// That;s the reason calling printCards() function which expects input of type 'deck'
	/*  This also works, anyway we call that
	newDeckFile := newDeckFromFile("my_decK_of_cards")
	newDeckFile.printCards()
	*/
	newDeckFromFile("my_decK_of_cards").printCards()

	// shuffle the deck and store shuffled data into deckOfCards()
	deckOfCards.shuffle()

	// now we use the shuffled deckOfCards to call printCards function again
	deckOfCards.printCards()
}
