package solitaire

import (
	"math"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
	"github.com/fyne-io/examples/solitaire/faces"
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
	return newTableRender(t)
}

// find card from an image, easier than keeping them in sync
func (t *Table) cardForPos(pos *canvas.Image) *Card {
	deck := NewSortedDeck()

	for i, face := range deck.Cards {
		if face.Face() == pos.Resource {
			card := NewCard((i%13)+1, Suit(math.Floor(float64(i)/13)))
			card.FaceUp = true // we know this as we checked the face
			return card
		}
	}

	return nil
}

func cardEquals(card1, card2 *Card) bool {
	if card1 == nil || card2 == nil {
		return card1 == nil && card2 == nil
	}

	return card1.Value == card2.Value && card1.Suit == card2.Suit
}

func (t *Table) cardTapped(cardPos *canvas.Image, pos fyne.Position, move func()) bool {
	if !withinCardBounds(cardPos, pos) {
		return false
	}

	card := t.cardForPos(cardPos)
	if cardPos.Resource != faces.ForSpace() && (card == nil || !card.FaceUp) {
		t.selected = nil
		widget.Renderer(t).(*tableRender).Refresh()

		return true
	}

	if t.selected == nil {
		t.selected = card
	} else {
		if !cardEquals(t.selected, card) {
			if move != nil {
				move()
			}
		}

		t.selected = nil
	}

	widget.Renderer(t).(*tableRender).Refresh()
	return true
}

func (t *Table) checkStackTapped(render *stackRender, stack *Stack, pos fyne.Position) bool {
	for i := len(stack.Cards) - 1; i >= 0; i-- {
		//		card := stack.Cards[i]

		if t.cardTapped(render.cards[i], pos, func() {
			t.game.MoveCardToStack(stack, t.selected)
		}) {
			return true
		}
	}

	return t.cardTapped(render.cards[0], pos, func() {
		t.game.MoveCardToStack(stack, t.selected)
	})
}

// OnMouseDown is called when the user taps the table widget
func (t *Table) OnMouseDown(event *fyne.MouseEvent) {
	render := widget.Renderer(t).(*tableRender)
	if withinCardBounds(render.deck, event.Position) {
		t.selected = nil
		t.game.DrawThree()
		render.Refresh()
		return
	}

	if t.game.Draw3 != nil {
		if t.cardTapped(render.pile3, event.Position, nil) {
			return
		}
	} else if t.game.Draw2 != nil {
		if t.cardTapped(render.pile2, event.Position, nil) {
			return
		}
	} else if t.game.Draw1 != nil {
		if t.cardTapped(render.pile1, event.Position, nil) {
			return
		}
	}

	if t.cardTapped(render.build1, event.Position, func() {
		t.game.MoveCardToBuild(t.game.Build1, t.selected)
	}) {
		return
	} else if t.cardTapped(render.build2, event.Position, func() {
		t.game.MoveCardToBuild(t.game.Build2, t.selected)
	}) {
		return
	} else if t.cardTapped(render.build3, event.Position, func() {
		t.game.MoveCardToBuild(t.game.Build3, t.selected)
	}) {
		return
	} else if t.cardTapped(render.build4, event.Position, func() {
		t.game.MoveCardToBuild(t.game.Build4, t.selected)
	}) {
		return
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

	t.selected = nil // clicked elsewhere
	widget.Renderer(t).(*tableRender).Refresh()
}

// NewTable creates a new table widget for the specified game
func NewTable(g *Game) *Table {
	return &Table{game: g}
}
