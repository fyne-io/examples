package solitaire

import "github.com/fyne-io/fyne"

type Stack struct {
	Cards []*Card
}

func (s *Stack) Push(card *Card) {
	s.Cards = append(s.Cards, card)
}

type Game struct {
	Deck Deck

	Draw1, Draw2, Draw3 *Card
	Drawn Deck

	Suit1 Stack
	Suit2 Stack
	Suit3 Stack
	Suit4 Stack

	Stack1 Stack
	Stack2 Stack
	Stack3 Stack
	Stack4 Stack
	Stack5 Stack
	Stack6 Stack
	Stack7 Stack
}

func pushToStack(s *Stack, d *Deck, count int) {
	for i := 0; i < count; i++ {
		s.Push(d.Pop())
	}
}

func (g *Game) Deal() {
	pushToStack(&g.Stack1, &g.Deck, 1)
	pushToStack(&g.Stack2, &g.Deck, 2)
	pushToStack(&g.Stack3, &g.Deck, 3)
	pushToStack(&g.Stack4, &g.Deck, 4)
	pushToStack(&g.Stack5, &g.Deck, 5)
	pushToStack(&g.Stack6, &g.Deck, 6)
	pushToStack(&g.Stack7, &g.Deck, 7)
}

func (g *Game) ResetDraw() {
	for ; len(g.Deck.Cards) > 0; g.DrawThree() {
	}

	// Reset the draw pile
	g.DrawThree()
}

func (g *Game) drawCard() *Card {
	if len(g.Deck.Cards) == 0 {
		return nil
	}

	popped := g.Deck.Pop()
	g.Drawn.Push(popped)
	return popped
}

func (g *Game) DrawThree() {
	if len(g.Deck.Cards) == 0 {
		g.Draw1 = nil
		g.Draw2 = nil
		g.Draw3 = nil

		g.Deck = g.Drawn
		g.Drawn = Deck{}
		return
	}

	g.Draw1 = g.drawCard()
	g.Draw2 = g.drawCard()
	g.Draw3 = g.drawCard()
}

func NewGame() *Game {
	game := &Game{}
	game.Deck = NewShuffledDeck()

	return game
}

func Show(app fyne.App) {
	game := NewGame()
	game.Deal()

	w := app.NewWindow("Solitaire")
	w.SetContent(NewTable(game))

	w.Show()
}
