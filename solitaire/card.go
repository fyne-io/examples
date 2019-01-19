package solitaire

import (
	"log"

	"fyne.io/fyne"

	"github.com/fyne-io/examples/solitaire/faces"
)

// Suit encodes one of the four possible suits for a playing card
type Suit int

// SuitColor represents the red/black of a suit
type SuitColor int

const (
	// SuitClubs is the "Clubs" playing card suit
	SuitClubs Suit = iota
	// SuitDiamonds is the "Diamonds" playing card suit
	SuitDiamonds
	// SuitHearts is the "Hearts" playing card suit
	SuitHearts
	// SuitSpades is the "Spades" playing card suit
	SuitSpades

	// SuitColorBlack is returned from Color() if the suit is Clubs or Spades
	SuitColorBlack SuitColor = iota
	// SuitColorRed is returned from Color() if the suit is Diamonds or Hearts
	SuitColorRed
)

const (
	// ValueJack is a convenience for the card 1 higher than 10
	ValueJack = 11
	// ValueQueen is the value for a queen face card
	ValueQueen = 12
	// ValueKing is the value for a king face card
	ValueKing = 13
)

// Card is a single playing card, it has a face value and a suit associated with it.
type Card struct {
	Value int
	Suit  Suit

	FaceUp bool
}

// Face returns a resource that can be used to render the associated card
func (c *Card) Face() fyne.Resource {
	return faces.ForCard(c.Value, int(c.Suit))
}

// TurnFaceUp sets the FaceUp field to true - so the card value can be seen
func (c *Card) TurnFaceUp() {
	c.FaceUp = true
}

// TurnFaceDown sets the FaceUp field to false - so the card should be hidden
func (c *Card) TurnFaceDown() {
	c.FaceUp = false
}

// Color returns the red or black color of the card suit
func (c *Card) Color() SuitColor {
	if c.Suit == SuitClubs || c.Suit == SuitSpades {
		return SuitColorBlack
	}

	return SuitColorRed
}

// NewCard returns a new card instance with the specified suit and value (1 based for Ace, 2 is 2 and so on).
func NewCard(value int, suit Suit) *Card {
	if value < 1 || value > 13 {
		log.Fatal("Invalid card face value")
	}

	return &Card{Value: value, Suit: suit}
}
