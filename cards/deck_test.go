package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expeceted deck length of 20, but got %v", len(d))
	}

	if d[0] != "spades of Ace" {
		t.Errorf("expected first card of Ace of Spedes, but got %v", d[0])
	}

	if d[len(d)-1] != "clubs of Four" {
		t.Errorf("expeceted last card of Four pf Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expeceted 16 cards in deck , got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
