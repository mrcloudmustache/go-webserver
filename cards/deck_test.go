package main

import (
	"os"
	"testing"
)

// t.Errorf("Expected deck length of 16, but got", len(d))
// to this:
//t.Errorf("Expected deck length of 16, but got %v", len(d))

func TestNewDeck(t *testing.T) {
	d := newDeck()
	length := 20
	if len(d) != length {
		t.Errorf("Expected deck length of %v, but got %v", length, len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card of %v, but got %v", "Ace of Clubs", d[0])
	}

	if d[len(d)-1] != "Four of Hearts" {
		t.Errorf("Expected last card of %v, but got %v", "Four of Hearts", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Delete any files with the new _decktesting
	os.Remove("_decktesting")

	// Create new deck and save testing file to disk
	d := newDeck()
	d.saveToFile("_decktesting")

	// Load testing file from disk
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 20 {
		t.Errorf("Expected deck length of %v, but got %v", 20, len(d))
	}

	// Remove temporary testing file
	os.Remove("_decktesting")
}
