package main

import (
	"os"
	"testing"
)

// Go treats any function starting with Test inside a _test.go file as a test
// The function name, by convention we are keeping as Test<FunctionName_we_want_to_test>
// t *testing.T is the testing context
func TestNewDeck(t *testing.T) {

	// Callig a real function we want to test, in this case we want to test newDeck() function of deck.go
	d := newDeck()

	// Test 01
	if len(d) != 16 {

		// Errorf is an one of the method an package called testing
		t.Errorf("Expected deck of length 16 but got %v", len(d))
	}

	// Test 02
	// When new deck is created fron newDeck() in deck.go, the first card in the slice should be Ace of Spades and we are making use of
	// that login for this test
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected First Card to be Ace of Spades but got %v", d[0])
	}

	// Test 03
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	// First remove "_decktesting" named file if present from previous test executions,
	//  just to make sure clean up is done before starting new tests
	_ = os.Remove("_decktesting")

	// Now, will test saveToFile function, by writing test case on that function
	deck := newDeck()

	// passing deck variable into saveToFile function
	deck.saveTofile("_decktesting")

	// Testing newDeckFromFile function, by passing an file
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in the deck, but got %v", len(loadedDeck))
	}

	// Again removing the file after test ran successfully..!!
	// Remove function return 2 values, 'nil' for successful execution and error for failures
	err := os.Remove("_decktesting")

	if err != nil {
		t.Fatalf("Failed to remove file after new test: %v", err)
	}

}
