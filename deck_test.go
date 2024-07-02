package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected 16 cards, got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, got %s", d[len(d)-1])
	}
}

func TestNewDeckFromFileAndSaveDeckInFile(t *testing.T) {
	os.Remove("_deck_test")
	deck := newDeck()
	deck.saveToFile("_deck_test")

	loadedDeck := newDeckFromFile("_deck_test")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards, got %d", len(loadedDeck))
	}

	os.Remove("_deck_test")
}

///t.Errorf("Expected deck length of 16, but got %v", len(d))
