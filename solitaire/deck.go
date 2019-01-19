package solitaire

import (
	"math/rand"
	"time"
)

// Deck is standard playing card collection, it contains up to 52 unique cards.
type Deck struct {
	Cards []*Card
}

// Shuffle reorganises the cards in the deck to a random order
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for c := 0; c < len(d.Cards); c++ {
		swap := rand.Intn(len(d.Cards))
		if swap != c {
			d.Cards[swap], d.Cards[c] = d.Cards[c], d.Cards[swap]
		}
	}
}

// Push adds the specified card to the top of the deck
func (d *Deck) Push(card *Card) {
	d.Cards = append(d.Cards, card)
}

// Pop removes the top card from the deck and returns it
func (d *Deck) Pop() *Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]

	return card
}

// Remove takes the specified card out of the deck
func (d *Deck) Remove(card *Card) {
	for i, c := range d.Cards {
		if cardEquals(c, card) {
			d.Cards = append(d.Cards[:i], d.Cards[i+1:]...)
		}
	}
}

// NewSortedDeck returns a standard deck in sorted order - starting with Ace of Clubs, ending with King of Spades.
func NewSortedDeck() *Deck {
	deck := &Deck{}

	c := 0
	suit := SuitClubs
	for i := 0; i < 4; i++ {
		for value := 1; value <= ValueKing; value++ {
			deck.Cards = append(deck.Cards, NewCard(value, suit))
			c++
		}
		suit++
	}

	return deck
}

// NewShuffledDeck returns a 52 card deck in random order
func NewShuffledDeck() *Deck {
	deck := NewSortedDeck()
	deck.Shuffle()

	return deck
}
