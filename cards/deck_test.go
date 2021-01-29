package main

import (
	"os"
	"testing"
)

// after new deck, validate length of deck
// validate that

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 20 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Six of Clubs" {
		t.Errorf("Expected Six of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	if err := os.Remove("_decktesting"); err != nil {
		t.Error(err)
	}
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but go %v", len(loadedDeck))
	}
	if err := os.Remove("_decktesting"); err != nil {
		t.Error(err)
	}
}
