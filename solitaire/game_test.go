package solitaire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestGame() *Game {
	return NewGameFromSeed(0xace)
}

func TestStack_Push(t *testing.T) {
	stack := &Stack{}
	card := NewCard(1, SuitSpades)

	assert.Equal(t, 0, len(stack.Cards))
	stack.Push(card)
	assert.Equal(t, 1, len(stack.Cards))
}

func TestGame_Deal(t *testing.T) {
	game := newTestGame()

	assert.Equal(t, 1, len(game.Stack1.Cards))
	assert.Equal(t, 2, len(game.Stack2.Cards))
	assert.Equal(t, 7, len(game.Stack7.Cards))
	assert.Equal(t, 24, len(game.Hand.Cards))
}

func TestGame_Deal_FaceUp(t *testing.T) {
	game := newTestGame()

	assert.Equal(t, true, game.Stack1.Cards[0].FaceUp)

	assert.Equal(t, false, game.Stack2.Cards[0].FaceUp)
	assert.Equal(t, true, game.Stack2.Cards[1].FaceUp)
}

func TestGame_Draw(t *testing.T) {
	game := newTestGame()

	assert.Equal(t, 24, len(game.Hand.Cards))

	game.DrawThree()
	assert.Equal(t, 21, len(game.Hand.Cards))
	assert.NotNil(t, game.Draw1)
	assert.True(t, game.Draw1.FaceUp)
	assert.NotNil(t, game.Draw2)
	assert.True(t, game.Draw2.FaceUp)
	assert.NotNil(t, game.Draw3)
	assert.True(t, game.Draw3.FaceUp)
}

func TestGameDrawEnd(t *testing.T) {
	game := newTestGame()

	assert.Equal(t, 24, len(game.Hand.Cards))

	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	game.DrawThree()
	assert.Equal(t, 0, len(game.Hand.Cards))
}

func TestGame_DrawCycles(t *testing.T) {
	game := newTestGame()

	assert.Equal(t, 24, len(game.Hand.Cards))

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
	assert.Equal(t, 24, len(game.Hand.Cards))
}

func TestGame_ResetDraw(t *testing.T) {
	game := newTestGame()

	assert.Equal(t, 24, len(game.Hand.Cards))

	game.DrawThree()
	game.ResetDraw()
	assert.Equal(t, 24, len(game.Hand.Cards))
}

func TestGame_DrawSymmetric(t *testing.T) {
	game := newTestGame()

	assert.Equal(t, 24, len(game.Hand.Cards))

	game.DrawThree()
	first := game.Draw1
	game.ResetDraw()

	// first draw again
	game.DrawThree()
	assert.Equal(t, first, game.Draw1)
}

func TestGame_MoveCardToBuildFromHand(t *testing.T) {
	game := newTestGame()
	game.DrawThree()
	game.Draw3.Value = 1

	game.MoveCardToBuild(game.Build1, game.Draw3)
	assert.Equal(t, 1, len(game.Build1.Cards))
	assert.Nil(t, game.Draw3)
}

func TestGame_MoveCardToBuildFromStack2(t *testing.T) {
	game := newTestGame()
	game.Stack2.Cards[1].Value = 1

	assert.Equal(t, 2, len(game.Stack2.Cards))
	assert.False(t, game.Stack2.Cards[0].FaceUp)

	game.MoveCardToBuild(game.Build2, game.Stack2.Cards[1])
	assert.Equal(t, 1, len(game.Build2.Cards))
	assert.Equal(t, 1, len(game.Stack2.Cards))
	assert.True(t, game.Stack2.Cards[0].FaceUp)
}

func TestGame_MoveCardToStack(t *testing.T) {
	game := newTestGame()

	game.Stack1.Cards[0].Value = 3
	game.Stack1.Cards[0].Suit = SuitClubs
	game.Stack2.Cards[1].Value = 2
	game.Stack2.Cards[1].Suit = SuitDiamonds
	assert.False(t, game.Stack2.Cards[0].FaceUp)

	game.MoveCardToStack(game.Stack1, game.Stack2.Cards[1])
	assert.Equal(t, 2, len(game.Stack1.Cards))
	assert.Equal(t, 1, len(game.Stack2.Cards))
	assert.True(t, game.Stack2.Cards[0].FaceUp)
}

func TestGame_MoveCardToStack_Empty(t *testing.T) {
	game := newTestGame()

	game.Stack1.Cards = []*Card{}
	game.Stack2.Cards[1].Value = ValueKing
	game.Stack2.Cards[1].Suit = SuitDiamonds

	game.MoveCardToStack(game.Stack1, game.Stack2.Cards[1])
	assert.Equal(t, 1, len(game.Stack1.Cards))
	assert.Equal(t, 1, len(game.Stack2.Cards))
	assert.True(t, game.Stack2.Cards[0].FaceUp)
}

func TestGame_MoveCardToStack_EmptyEmpty(t *testing.T) {
	game := newTestGame()

	game.Stack3.Cards = []*Card{}
	king := NewCard(ValueKing, SuitDiamonds)
	game.Stack2.Cards = []*Card{king}

	game.MoveCardToStack(game.Stack3, game.Stack2.Cards[0])
	assert.Equal(t, 1, len(game.Stack3.Cards))
	assert.Equal(t, 0, len(game.Stack2.Cards))
}

func TestGame_MoveCardToStack_Stack(t *testing.T) {
	game := newTestGame()

	game.Stack1.Cards[0].Value = 7
	game.Stack1.Cards[0].Suit = SuitClubs
	game.Stack2.Cards[0].Value = 6
	game.Stack2.Cards[0].Suit = SuitDiamonds
	game.Stack2.Cards[1].Value = 5
	game.Stack2.Cards[1].Suit = SuitSpades

	game.MoveCardToStack(game.Stack1, game.Stack2.Cards[0])
	assert.Equal(t, 3, len(game.Stack1.Cards))
	assert.Equal(t, 0, len(game.Stack2.Cards))
}

func TestGame_MoveCardToStack_KingStack(t *testing.T) {
	game := newTestGame()

	game.Stack1.Cards = []*Card{}
	game.Stack3.Cards[1].Value = ValueKing
	game.Stack3.Cards[1].Suit = SuitDiamonds
	game.Stack3.Cards[2].Value = ValueQueen
	game.Stack3.Cards[2].Suit = SuitSpades

	game.MoveCardToStack(game.Stack1, game.Stack3.Cards[1])
	assert.Equal(t, 2, len(game.Stack1.Cards))
	assert.Equal(t, 1, len(game.Stack3.Cards))
}
