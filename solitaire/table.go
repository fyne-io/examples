package solitaire

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/widget"
)

// Table represents the rendering of a game in progress
type Table struct {
	size     fyne.Size
	position fyne.Position
	hidden   bool

	game *Game
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

func withinBounds(pos fyne.Position, card *canvas.Image) bool {
	if pos.X < card.Position().X || pos.Y < card.Position().Y {
		return false
	}

	if pos.X >= card.Position().X+card.Size().Width || pos.Y >= card.Position().Y+card.Size().Height {
		return false
	}

	return true
}

// OnMouseDown is called when the user taps the table widget
func (t *Table) OnMouseDown(event *fyne.MouseEvent) {
	if withinBounds(event.Position, widget.Renderer(t).(*tableRender).deck) {
		t.game.DrawThree()
		widget.Renderer(t).Refresh()
	}
}

// NewTable creates a new table widget for the specified game
func NewTable(g *Game) *Table {
	return &Table{game: g}
}
