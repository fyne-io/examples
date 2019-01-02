package solitaire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	stack := &Stack{}
	card := NewCard(1, SuitSpades)

	assert.Equal(t, 0, len(stack.Cards))
	stack.Push(card)
	assert.Equal(t, 1, len(stack.Cards))
}

func TestGame_Deal(t *testing.T) {
	game := NewGame()

	assert.Equal(t, 1, len(game.Stack1.Cards))
	assert.Equal(t, 2, len(game.Stack2.Cards))
	assert.Equal(t, 7, len(game.Stack7.Cards))
	assert.Equal(t, 24, len(game.Deck.Cards))
}

func TestGame_Draw(t *testing.T) {
	game := NewGame()

	assert.Equal(t, 24, len(game.Deck.Cards))

	game.DrawThree()
	assert.Equal(t, 21, len(game.Deck.Cards))
	assert.NotNil(t, game.Draw1)
	assert.NotNil(t, game.Draw2)
	assert.NotNil(t, game.Draw3)
}

func TestGameDrawEnd(t *testing.T) {
	game := NewGame()

	assert.Equal(t, 24, len(game.Deck.Cards))

	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	assert.Equal(t, 0, len(game.Deck.Cards))
}

func TestGame_DrawCycles(t *testing.T) {
	game := NewGame()

	assert.Equal(t, 24, len(game.Deck.Cards))

	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()

	// This is the extra one...
	game.DrawThree()
	assert.Equal(t, 24, len(game.Deck.Cards))
}

func TestGame_ResetDraw(t *testing.T) {
	game := NewGame()

	assert.Equal(t, 24, len(game.Deck.Cards))

	game.DrawThree()
	game.ResetDraw()
	assert.Equal(t, 24, len(game.Deck.Cards))
}

func TestGame_DrawSymmetric(t *testing.T) {
	game := NewGame()

	assert.Equal(t, 24, len(game.Deck.Cards))

	game.DrawThree()
	first := game.Draw1
	game.ResetDraw()

	// first draw again
	game.DrawThree()
	assert.Equal(t, first, game.Draw1)
}
