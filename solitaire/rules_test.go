package solitaire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleCanMoveToBuild_Empty(t *testing.T) {
	g := NewGame()
	card := NewCard(1, SuitClubs)

	assert.True(t, g.ruleCanMoveToBuild(g.Build1, card))
	card.Value = 3
	assert.False(t, g.ruleCanMoveToBuild(g.Build1, card))
}

func TestRuleCanMoveToBuild_Over(t *testing.T) {
	g := NewGame()
	card := NewCard(1, SuitClubs)
	g.Build1.Push(card)

	card = NewCard(2, SuitClubs)
	assert.True(t, g.ruleCanMoveToBuild(g.Build1, card))
	card.Suit = SuitDiamonds
	assert.False(t, g.ruleCanMoveToBuild(g.Build1, card))
}
