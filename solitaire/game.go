package solitaire

import "fyne.io/fyne"

// Stack represents a number of cards in a particular order
type Stack struct {
	Cards []*Card
}

// Push adds a new card to the top of the stack
func (s *Stack) Push(card *Card) {
	s.Cards = append(s.Cards, card)
}

// Top gets the top card of a stack, or nil if the stack is empty
func (s *Stack) Top() *Card {
	if len(s.Cards) == 0 {
		return nil
	}

	return s.Cards[len(s.Cards)-1]
}

// Pop removes the top card of a stack, turning up any card immediately underneath and returning the removed card
func (s *Stack) Pop() *Card {
	if len(s.Cards) == 0 {
		return nil
	}

	ret := s.Top()
	s.Cards = s.Cards[0 : len(s.Cards)-1]

	if len(s.Cards) > 0 {
		s.Cards[len(s.Cards)-1].TurnFaceUp()
	}
	return ret
}

// Game represents a full solitaire game, starting from a standard draw
type Game struct {
	Hand *Deck

	Draw1, Draw2, Draw3 *Card
	Drawn               *Deck

	Build1 *Stack
	Build2 *Stack
	Build3 *Stack
	Build4 *Stack

	Stack1 *Stack
	Stack2 *Stack
	Stack3 *Stack
	Stack4 *Stack
	Stack5 *Stack
	Stack6 *Stack
	Stack7 *Stack
}

func pushToStack(s *Stack, d *Deck, count int) {
	for i := 0; i < count; i++ {
		card := d.Pop()
		if i == count-1 {
			card.FaceUp = true
		}
		s.Push(card)
	}
}

func (g *Game) deal() {
	pushToStack(g.Stack1, g.Hand, 1)
	pushToStack(g.Stack2, g.Hand, 2)
	pushToStack(g.Stack3, g.Hand, 3)
	pushToStack(g.Stack4, g.Hand, 4)
	pushToStack(g.Stack5, g.Hand, 5)
	pushToStack(g.Stack6, g.Hand, 6)
	pushToStack(g.Stack7, g.Hand, 7)
}

// ResetDraw resets the draw pile to be completely available (no cards drawn)
func (g *Game) ResetDraw() {
	for ; len(g.Hand.Cards) > 0; g.DrawThree() {
	}

	// Reset the draw pile
	g.DrawThree()
}

func (g *Game) drawCard() *Card {
	if len(g.Hand.Cards) == 0 {
		return nil
	}

	popped := g.Hand.Pop()
	popped.FaceUp = true
	g.Drawn.Push(popped)
	return popped
}

// DrawThree draws three cards from the deck and adds them to the draw pile(s).
// If there are no cards available to be drawn it will cycle back to the beginning and draw the first three.
func (g *Game) DrawThree() {
	if len(g.Hand.Cards) == 0 {
		g.Draw1 = nil
		g.Draw2 = nil
		g.Draw3 = nil

		g.Hand = g.Drawn
		for _, card := range g.Hand.Cards {
			card.TurnFaceDown()
		}
		g.Drawn = &Deck{}
		return
	}

	g.Draw1 = g.drawCard()
	g.Draw2 = g.drawCard()
	g.Draw3 = g.drawCard()
}

// MoveCardToBuild attempts to move the currently selected card to a build stack.
// If the move is not possible it will return.
func (g *Game) MoveCardToBuild(build *Stack, card *Card) {
	if !g.ruleCanMoveToBuild(build, card) {
		return
	}

	g.removeCard(card)
	build.Push(card)
}

// MoveCardToStack attempts to move the currently selected card to a table stack.
// If the move is not possible it will return.
func (g *Game) MoveCardToStack(stack *Stack, card *Card) {
	// TODO if we can then move - and everything under the card too...
}

func (g *Game) removeCard(card *Card) {
	if cardEquals(card, g.Draw3) {
		g.Drawn.Remove(card)
		g.Draw3 = nil
	} else if cardEquals(card, g.Draw2) {
		g.Drawn.Remove(card)
		g.Draw2 = nil
	} else if cardEquals(card, g.Draw2) {
		g.Drawn.Remove(card)
		g.Draw2 = nil
	} else if cardEquals(card, g.Stack1.Top()) {
		g.Stack1.Pop()
	} else if cardEquals(card, g.Stack2.Top()) {
		g.Stack2.Pop()
	} else if cardEquals(card, g.Stack3.Top()) {
		g.Stack3.Pop()
	} else if cardEquals(card, g.Stack4.Top()) {
		g.Stack4.Pop()
	} else if cardEquals(card, g.Stack5.Top()) {
		g.Stack5.Pop()
	} else if cardEquals(card, g.Stack6.Top()) {
		g.Stack6.Pop()
	} else if cardEquals(card, g.Stack7.Top()) {
		g.Stack7.Pop()
	}
}

// NewGame starts a new solitaire game and draws to the standard configuration
func NewGame() *Game {
	game := &Game{}
	game.Hand = NewShuffledDeck()

	game.Drawn = &Deck{}

	game.Stack1 = &Stack{}
	game.Stack2 = &Stack{}
	game.Stack3 = &Stack{}
	game.Stack4 = &Stack{}
	game.Stack5 = &Stack{}
	game.Stack6 = &Stack{}
	game.Stack7 = &Stack{}

	game.Build1 = &Stack{}
	game.Build2 = &Stack{}
	game.Build3 = &Stack{}
	game.Build4 = &Stack{}

	game.deal()
	return game
}

// Show creates a new game and loads a table rendered in a new window.
func Show(app fyne.App) {
	game := NewGame()

	w := app.NewWindow("Solitaire")
	w.SetPadded(false)
	w.SetContent(NewTable(game))

	w.Show()
}
