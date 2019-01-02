package solitaire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertNotEqualCard(t *testing.T, value int, suit Suit, card *Card) {
	if value != card.Value || suit != card.Suit {
		return
	}

	t.Fail()
}

func TestNewDeck(t *testing.T) {
	deck := NewSortedDeck()

	assert.Equal(t, 52, len(deck.Cards))
}

func TestNewSortedDeck(t *testing.T) {
	deck := NewSortedDeck()

	assert.Equal(t, 1, deck.Cards[0].Value)
	assert.Equal(t, SuitClubs, deck.Cards[0].Suit)

	assert.Equal(t, 13, deck.Cards[12].Value)
	assert.Equal(t, SuitClubs, deck.Cards[0].Suit)

	assert.Equal(t, 1, deck.Cards[13].Value)
	assert.Equal(t, SuitDiamonds, deck.Cards[13].Suit)

	assert.Equal(t, 5, deck.Cards[30].Value)
	assert.Equal(t, SuitHearts, deck.Cards[30].Suit)

	assert.Equal(t, 11, deck.Cards[49].Value)
	assert.Equal(t, SuitSpades, deck.Cards[49].Suit)
}

func TestNewShuffledDeck(t *testing.T) {
	deck := NewShuffledDeck()

	// TODO make this more predictable!
	assertNotEqualCard(t, 1, SuitClubs, deck.Cards[0])
	assertNotEqualCard(t, 13, SuitClubs, deck.Cards[12])
	assertNotEqualCard(t, 1, SuitDiamonds, deck.Cards[13])
}

func TestDeck_Push(t *testing.T) {
	deck := Deck{}

	assert.Equal(t, 0, len(deck.Cards))
	card := NewCard(1, SuitDiamonds)
	deck.Push(card)

	assert.Equal(t, 1, len(deck.Cards))
	assert.Equal(t, card, deck.Pop())
}

func TestDeck_Pop(t *testing.T) {
	deck := NewSortedDeck()

	assert.Equal(t, 52, len(deck.Cards))
	card := deck.Pop()

	assert.NotNil(t, card)
	assert.Equal(t, 51, len(deck.Cards))
}
