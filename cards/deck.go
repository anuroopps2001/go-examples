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

import (
	"fmt"
	"math/rand"
	"os"      // to conver single into byte slice using WriteFile function of this package and generate a file
	"strings" // importing strings package to join slice of strings into single string separated by ',' as per our requirement
)

// custom types
// creating a new custom type called deck which basically a slice of strings
// think of this like lable i.e deck = []string
type deck []string

// creating new function with return type of slice i,e []string and because return type is used with the function,
//
//	we are using return at the end of the function and Println
//
// []string == deck custom type here
func newDeck() deck {
	cards := deck{} // creating empty slice of type deck

	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards // returning the variable declared inside the function and also because we used retun type at function declaration
}

// function with variable name called cards and its return type i.e slice or written as []string and return type specified as int

func (d deck) countCards() int {
	return len(d)
}

/* func lastCard(cards deck) string {
	if len(cards) == 0 {
		return "No cards available"
	} else {
		return cards[len(cards)-1]
	}
} */

func (d deck) lastCard() string {
	if len(d) == 0 {
		return "No Cards Available"
	} else {
		return d[len(d)-1]
	}
}

/*
	 func firstCard(cards deck) string {
		if len(cards) == 0 {
			return "No cards available"
		} else {
			return cards[0]
		}

}
*/
func (d deck) firstCard() string {
	if len(d) == 0 {
		return "No Cards Available"
	} else {
		return d[0]
	}
}

/* func printCards(cards deck) {
	for _, card := range cards {
		fmt.Println(card)
	}
} */

func (d deck) printCards() { // call this function with any variable that's holding values of type of 'deck'
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func <receiver i,e (d deck) in this example> <function_name> {}  -- syntax while using receivers inside go functions

// 'd' is a receiver and it;s of custom type 'deck' which is equal to slice of strings []string
// calling functions involved with receivers -

/*
// cards := deck{"a", "b", "c"}  -- cards is a variable of type deck
// func (d deck) printCards(){   -- using receiver of type deck i.e 'd'  and cards variable is equal to receiver d i,e 'cards == d'
//   for i, card := range d {
//     fmt.println(card)}
//}
// cards.printCards() -- calling printCards function with cards variable, because the variable is of type deck */

/*
cards is a variable

its type is deck */

/*
Methods are attached to types.
Any value of that type can call those methods.

In our example, printCards is a method of type deck and it was called using the variable which also the values of type deck


cards.PrintCards() -- now the value of cards variable will be assigned to receiver we used with printCards method i.e 'd'
and now inside PrintCards function, cards variable will be replaced by receiver 'd'

That's why inside printCards() function, under for loop, we used "for i, card := range d" and not "for i, card := range cards"

because now cards variable replaced by receiver defined with 'd'
*/

func dealOfDeck(d deck, handSize int) (deck, deck) { // we should call this function with first arg being type deck and second arg being type int
	// and this function returns 2 values of both of type deck and that's what (deck, deck) mean at the function declaration

	return d[:handSize], d[handSize:]
}

// converting deck into string using byte slice concept
func (d deck) toString() string {
	// []string(d) //byte slicing and it will return slice of strings by taking deck type as an input
	// also created a single string from output of deck byte slicing with join function seprated by ','
	return strings.Join([]string(d), ",")
}

// Based on the function call, we need to decide the function signature, i,e whereher we need to used receiver or not
// check out, function of this specific function to know more, why we have used receiver here
// function to take single string above function and write into the file on local system using WriteFile function of os package
func (d deck) saveTofile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Based on the function call, we need to decide the function signature, i,e whereher we need to used receiver or not
// check out, function of this specific function to know more, why we have not used receiver here
// Readfile expects filename as an arg which is of string type and returns a value of type deck in our case
func newDeckFromFile(filename string) deck {

	bs, err := os.ReadFile(filename) // Reafile returns 2 values, i.e []byte and error and here, we are storing []byte as bs and error as err

	// If everything goes well with this function, then err returns 'nil' and if something goes wrong with this functions. then the return value wil; not be equal to nil
	// and we are handling that error here with if condition
	if err != nil {
		fmt.Println("Error: ", err) // print the error
		os.Exit(1)                  // non-zero exit code, unsuccessful and exiting the program
	}

	// err == nil, that means, it will return byte slice and we will convert byte slice into single string using type conversion

	// type conversion syntax: what_we_want(what)we_have) and converting above generated new byte slice into single string
	// and storing into variable called s which is a single string
	s := string(bs) // s stands for single string, holding all the values

	// converting a single string holding all the values into multiple strings separated by ','
	ms := strings.Split(s, ",") // ms stands for multiple strings separated by ','

	// type conversion, converting strings into deck type
	return deck(ms)
}

func (d deck) shuffle() {
	for index := range d {
		// rand function will generate a random between 0 and length-1 of deck type
		newPosition := rand.Intn(len(d) - 1)

		// swap index with newPosition and newPosition with index
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
