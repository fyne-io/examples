package solitaire

import (
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/canvas"
)

// Table represents the rendering of a game in progress
type Table struct {
	size     fyne.Size
	position fyne.Position
	hidden   bool

	game     *Game
	renderer *tableRender
}

// CurrentSize gets the current size of the table
func (t *Table) CurrentSize() fyne.Size {
	return t.size
}

// Resize sets the current size of the table
func (t *Table) Resize(size fyne.Size) {
	t.size = size
	t.Renderer().Layout(size)
}

// CurrentPosition gets the current position of the table
func (t *Table) CurrentPosition() fyne.Position {
	return t.position
}

// Move sets the current position of the table
func (t *Table) Move(pos fyne.Position) {
	t.position = pos
	t.Renderer().Layout(t.size)
}

// MinSize specifies the minimum size of a table
func (t *Table) MinSize() fyne.Size {
	return t.Renderer().MinSize()
}

// IsVisible returns true if the table widget is currently visible
func (t *Table) IsVisible() bool {
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
	t.Renderer().ApplyTheme()
}

// Renderer gets the widget renderer for this table - internal use only
func (t *Table) Renderer() fyne.WidgetRenderer {
	if t.renderer == nil {
		t.renderer = newTableRender(t.game)
	}

	t.renderer.Refresh()
	return t.renderer
}

func withinBounds(pos fyne.Position, card *canvas.Image) bool {
	if pos.X < card.Position.X || pos.Y < card.Position.Y {
		return false
	}

	if pos.X >= card.Position.X+card.Size.Width || pos.Y >= card.Position.Y+card.Size.Height {
		return false
	}

	return true
}

// OnMouseDown is called when the user taps the table widget
func (t *Table) OnMouseDown(event *fyne.MouseEvent) {
	if withinBounds(event.Position, t.renderer.deck) {
		t.game.DrawThree()
		t.Renderer().Refresh()
	}
}

// NewTable creates a new table widget for the specified game
func NewTable(g *Game) *Table {
	return &Table{game: g}
}
