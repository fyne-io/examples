package solitaire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	card := NewCard(3, SuitClubs)

	assert.False(t, card.FaceUp)
}

func TestCard_TurnFaceUp(t *testing.T) {
	card := NewCard(3, SuitClubs)
	card.TurnFaceUp()

	assert.True(t, card.FaceUp)
}

func TestCard_TurnFaceDown(t *testing.T) {
	card := NewCard(3, SuitClubs)
	card.TurnFaceUp()
	card.TurnFaceDown()

	assert.False(t, card.FaceUp)
}
