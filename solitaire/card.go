package solitaire

import (
	"github.com/fyne-io/fyne"
	"log"
)
import "github.com/fyne-io/examples/solitaire/faces"

type Suit int

const (
	ValueJack  = 11
	ValueQueen = 12
	ValueKing  = 13
)

const (
	SuitClubs Suit = iota
	SuitDiamonds
	SuitHearts
	SuitSpades
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
