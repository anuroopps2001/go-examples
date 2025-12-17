package main

import "testing"

// The function name, by convention we are keeping as Test<FunctionName_we_want_to_test>
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16 but got %v", len(d))
	}
}
