package solitaire

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := &Stack{}
	card := NewCard(1, SUIT_SPADES)

	assert.Equal(t, 0, len(stack.Cards))
	stack.Push(card)
	assert.Equal(t, 1, len(stack.Cards))
}

func TestNewGame(t *testing.T) {
	game := NewGame()

	assert.Equal(t, 52, len(game.Deck.Cards))
}

func TestGame_Deal(t *testing.T) {
	game := NewGame()
	game.Deal()

	assert.Equal(t, 1, len(game.Stack1.Cards))
	assert.Equal(t, 2, len(game.Stack2.Cards))
	assert.Equal(t, 7, len(game.Stack7.Cards))
	assert.Equal(t, 24, len(game.Deck.Cards))
}

func TestGame_Draw(t *testing.T) {
	game := NewGame()
	game.Deal()

	assert.Equal(t, 24, len(game.Deck.Cards))

	game.Draw()
	assert.Equal(t, 21, len(game.Deck.Cards))
	assert.NotNil(t, game.Draw1)
	assert.NotNil(t, game.Draw2)
	assert.NotNil(t, game.Draw3)
}

func TestGameDrawEnd(t *testing.T) {
	game := NewGame()
	game.Deal()

	assert.Equal(t, 24, len(game.Deck.Cards))

	game.Draw()
	game.Draw()
	game.Draw()
	game.Draw()
	game.Draw()
	game.Draw()
	game.Draw()
	game.Draw()
	assert.Equal(t, 0, len(game.Deck.Cards))

	game.Draw() // this could crash
}
