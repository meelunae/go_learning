package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[51] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[51])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	os.Remove(".decktesting")

	deck := newDeck()
	deck.saveToFile(".decktesting")
	deckFromFile := loadDeckFromFile(".decktesting")

	if len(deckFromFile) != 52 {
		t.Errorf("Expected length of 52, but got %v", len(deckFromFile))
	}

	if deckFromFile[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", deckFromFile[0])
	}

	if deckFromFile[51] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", deckFromFile[51])
	}

	os.Remove(".decktesting")
}
