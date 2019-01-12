package solitaire

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// Table represents the rendering of a game in progress
type Table struct {
	size     fyne.Size
	position fyne.Position
	hidden   bool

	game     *Game
	selected *Card
}

// Size gets the current size of the table
func (t *Table) Size() fyne.Size {
	return t.size
}

// Resize sets the current size of the table
func (t *Table) Resize(size fyne.Size) {
	t.size = size
	widget.Renderer(t).Layout(size)
}

// Position gets the current position of the table
func (t *Table) Position() fyne.Position {
	return t.position
}

// Move sets the current position of the table
func (t *Table) Move(pos fyne.Position) {
	t.position = pos
	widget.Renderer(t).Layout(t.size)
}

// MinSize specifies the minimum size of a table
func (t *Table) MinSize() fyne.Size {
	return widget.Renderer(t).MinSize()
}

// Visible returns true if the table widget is currently visible
func (t *Table) Visible() bool {
	return !t.hidden
}

// Show sets the table widget to be visible
func (t *Table) Show() {
	t.hidden = false
}

// Hide sets the table widget to be hidden
func (t *Table) Hide() {
	t.hidden = true
}

// ApplyTheme updates the widget with the current theme
func (t *Table) ApplyTheme() {
	widget.Renderer(t).ApplyTheme()
}

// CreateRenderer gets the widget renderer for this table - internal use only
func (t *Table) CreateRenderer() fyne.WidgetRenderer {
	return newTableRender(t.game)
}

func (t *Table) selectCard(card *cardPosition) {
	render := widget.Renderer(t).(*tableRender)
	if card == nil || t.selected == card.card {
		t.selected = nil
		render.selectCard(nil)
		return // don'd reselect the same card
	}

	t.selected = card.card
	render.selectCard(card)
}

func (t *Table) cardTapped(card *cardPosition, pos fyne.Position) bool {
	if !card.card.FaceUp {
		return false
	}

	if card.withinBounds(pos) {
		t.selectCard(card)
		return true
	}

	return false
}

func (t *Table) checkStackTapped(render *stackRender, stack Stack, pos fyne.Position) bool {
	for i := len(stack.Cards) - 1; i >= 0; i-- {
		card := stack.Cards[i]
		if !card.FaceUp {
			return false
		}

		if t.cardTapped(render.cards[i], pos) {
			return true
		}
	}

	return false
}

// OnMouseDown is called when the user taps the table widget
func (t *Table) OnMouseDown(event *fyne.MouseEvent) {
	render := widget.Renderer(t).(*tableRender)
	if render.deck.withinBounds(event.Position) {
		t.selectCard(nil)
		t.game.DrawThree()

		// TODO consier moving this
		render.pile1.card = t.game.Draw1
		render.pile2.card = t.game.Draw2
		render.pile3.card = t.game.Draw3
		render.Refresh()
		return
	}

	// TODO check this!
	if t.game.Draw3 != nil {
		if t.cardTapped(render.pile3, event.Position) {
			return
		}
	} else if t.game.Draw2 != nil {
		if t.cardTapped(render.pile2, event.Position) {
			return
		}
	} else if t.game.Draw1 != nil {
		if t.cardTapped(render.pile1, event.Position) {
			return
		}
	}

	if t.checkStackTapped(render.stack1, t.game.Stack1, event.Position) {
		return
	} else if t.checkStackTapped(render.stack2, t.game.Stack2, event.Position) {
		return
	} else if t.checkStackTapped(render.stack3, t.game.Stack3, event.Position) {
		return
	} else if t.checkStackTapped(render.stack4, t.game.Stack4, event.Position) {
		return
	} else if t.checkStackTapped(render.stack5, t.game.Stack5, event.Position) {
		return
	} else if t.checkStackTapped(render.stack6, t.game.Stack6, event.Position) {
		return
	} else if t.checkStackTapped(render.stack7, t.game.Stack7, event.Position) {
		return
	}

	t.selectCard(nil) // clicked elsewhere
}

// NewTable creates a new table widget for the specified game
func NewTable(g *Game) *Table {
	return &Table{game: g}
}
