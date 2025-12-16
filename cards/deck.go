/* package main

import "fmt"

// creating new custom type of "deck" which is a slice of string i.e []string
type deck []string // custom type deck of real type []string

// creating new deck
func newDeck() deck {
	cards := deck{} // Empty slice i.e  []string{}

	// created slice of suites
	cardSuites := []string{"Spades", "Diamonds", "Hearts"}

	// created slice of values
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ (underscore) says that, these are variables defined but we are not using them anywhere
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) printCards() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
*/

package main

import "fmt"

// custom types

// think of this like lable i.e deck = []string
type deck []string

// creatina new function with return type of slice i,e []string and because return type is used with the function,
//  we are using return at the end of the function and Println
func newDeck() deck {
	cards := deck{}

	suits := deck{"Spades", "Diamonds", "Hearts"}
	values := deck{"Ace", "Two", "Three", "Four"}

	for _, suite := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards

}

// function with variable name called cards and its return type i.e slice or written as []string and return type specified as int
func countCards(cards deck) int {
	return len(cards)
}

func lastCard(cards deck) string {
	if len(cards) == 0 {
		return "No cards available"
	} else {
		return cards[len(cards)-1]
	}
}

func firstCard(cards deck) string {
	if len(cards) == 0 {
		return "No cards available"
	} else {
		return cards[0]
	}

}

func printCards(cards deck) {
	for _, card := range cards {
		fmt.Println(card)
	}
}
