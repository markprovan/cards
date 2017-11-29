package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	assert.NotNil(t, d)
	assert.Equal(t, 16, len(d))
	assert.Equal(t, "Ace of Spades", d[0])
	assert.Equal(t, "Four of Clubs", d[len(d)-1])
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	assert.NotNil(t, loadedDeck)
	assert.Equal(t, 16, len(loadedDeck))
	assert.Equal(t, "Ace of Spades", loadedDeck[0])
	assert.Equal(t, "Four of Clubs", loadedDeck[len(loadedDeck)-1])

	os.Remove("_decktesting")
}
