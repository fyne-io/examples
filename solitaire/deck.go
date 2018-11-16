package solitaire

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []*Card
}

func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for c := 0; c < len(d.Cards); c++ {
		swap := rand.Intn(len(d.Cards))
		if swap != c {
			d.Cards[swap], d.Cards[c] = d.Cards[c], d.Cards[swap]
		}
	}
}

func (d *Deck) Push(card *Card) {
	d.Cards = append(d.Cards, card)
}

func (d *Deck) Pop() *Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card
}

func NewSortedDeck() Deck {
	deck := Deck{}

	c := 0
	suit := SUIT_CLUBS
	for i := 0; i < 4; i++ {
		for value := 1; value <= VALUE_KING; value++ {
			deck.Cards = append(deck.Cards, NewCard(value, suit))
			c++
		}
		suit++
	}

	return deck
}

func NewShuffledDeck() Deck {
	deck := NewSortedDeck()
	deck.Shuffle()

	return deck
}
