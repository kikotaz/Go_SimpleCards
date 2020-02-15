package main

import (
	"os"
	"testing"
)

//Testingest the newDeck() function
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52, but found %v", len(d))
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card King of Clubs, but found %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, got %v ", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
