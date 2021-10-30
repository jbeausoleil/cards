package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()

	// Write statement to see if the deck has the right number of cards
	if len(d) != 52 {
		// If it doesn't, tell the go test handler that there is something wrong
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// Assert that the first card in the deck is "Ace of Spades", and last card is "King of Clubs"
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	errRemove := os.Remove("_decktesting")
	if errRemove != nil {
		return
	} // remove any crashed files from the past

	// Create deck and write to file
	deck := newDeck()
	errSave := deck.saveToFile("_decktesting")
	if errSave != nil {
		return 
	}

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in the deck, but got %v", len(loadedDeck))
	}
}