package solitaire

import (
	"github.com/fyne-io/fyne"
	"log"
)
import "github.com/fyne-io/examples/solitaire/faces"

type Suit int

const (
	VALUE_JACK  = 11
	VALUE_QUEEN = 12
	VALUE_KING  = 13
)

const (
	SUIT_CLUBS Suit = iota
	SUIT_DIAMONDS
	SUIT_HEARTS
	SUIT_SPADES
)

type Card struct {
	Value int
	Suit  Suit
}

func (c Card) Face() fyne.Resource {
	return faces.ForCard(c.Value, int(c.Suit))
}

func NewCard(value int, suit Suit) *Card {
	if value < 1 || value > 13 {
		log.Fatal("Invalid card face value")
	}

	return &Card{value, suit}
}
