package solitaire

import (
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/canvas"
)

type table struct {
	size     fyne.Size
	position fyne.Position

	game     *Game
	renderer *tableRender
}

func (t *table) CurrentSize() fyne.Size {
	return t.size
}

func (t *table) Resize(size fyne.Size) {
	t.size = size
	t.Renderer().Layout(size)
}

func (t *table) CurrentPosition() fyne.Position {
	return t.position
}

func (t *table) Move(pos fyne.Position) {
	t.position = pos
	t.Renderer().Layout(t.size)
}

func (t *table) MinSize() fyne.Size {
	return t.Renderer().MinSize()
}

func (t *table) IsVisible() bool {
	return t.hidden
}

func (t *table) Show() {
	t.hidden = false
}

func (t *table) Hide() {
	t.hidden = true
}

func (t *table) ApplyTheme() {
	t.Renderer().ApplyTheme()
}

func (t *table) Renderer() fyne.WidgetRenderer {
	if t.renderer == nil {
		t.renderer = newTableRender(t.game)
	}

	t.renderer.Refresh()
	return t.renderer
}

func (t *table) OnMouseDown(event *fyne.MouseEvent) {
	if event.Position.X <= cardSize.Width+smallPad && event.Position.Y <= cardSize.Height+smallPad {
		t.game.Draw()
		t.Renderer().Refresh()
	}
}

func NewTable(g *Game) *table {
	return &table{game: g}
}
