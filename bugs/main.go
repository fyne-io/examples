package bugs

import (
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/dialog"
	"github.com/fyne-io/fyne/layout"
	"github.com/fyne-io/fyne/widget"
)

type game struct {
	board *board

	size     fyne.Size
	position fyne.Position

	window   fyne.Window
	renderer *gameRenderer
}

func (g *game) CurrentSize() fyne.Size {
	return g.size
}

func (g *game) Resize(size fyne.Size) {
	g.size = size
	g.Renderer().Layout(size)
}

func (g *game) CurrentPosition() fyne.Position {
	return g.position
}

func (g *game) Move(pos fyne.Position) {
	g.position = pos
	g.Renderer().Layout(g.size)
}

func (g *game) MinSize() fyne.Size {
	return g.Renderer().MinSize()
}

func (g *game) ApplyTheme() {
	g.Renderer().ApplyTheme()
}

func (g *game) Renderer() fyne.WidgetRenderer {
	if g.renderer == nil {
		g.renderer = g.createRenderer()
	}

	return g.renderer
}

type gameRenderer struct {
	grid *fyne.Container

	game *game
}

func (g *gameRenderer) MinSize() fyne.Size {
	return fyne.NewSize(g.game.board.width*36, g.game.board.height*36)
}

func (g *gameRenderer) Layout(size fyne.Size) {
	g.grid.Layout.Layout(g.grid.Objects, size)
}

func (g *gameRenderer) ApplyTheme() {
}

func (g *gameRenderer) Refresh() {
	fyne.RefreshObject(g.grid)
}

func (g *gameRenderer) Objects() []fyne.CanvasObject {
	return g.grid.Objects
}

func (g *game) refreshSquare(x, y int) {
	if x < 0 || y < 0 || x >= g.board.width || y >= g.board.height {
		return
	}

	sq := g.board.bugs[y][x]
	i := y*g.board.width + x
	button := g.renderer.grid.Objects[i].(*widget.Button)

	if !sq.shown {
		if button.Icon == code {
			return
		}
		button.Icon = code
		button.Text = ""
	} else if sq.bug {
		if button.Icon == bug {
			return
		}
		button.Icon = bug
		button.Text = ""
	} else if button.Icon == nil {
		return
	} else {
		button.Icon = nil
		button.Text = squareString(sq)
	}

	// avoid double refresh that Setxxx would cause
	button.Renderer().Refresh()
}

func (g *game) refreshAround(xp, yp, d int) {
	x, y := xp-d, yp-d
	for ; x < xp+d; x++ {
		g.refreshSquare(x, y)
	}
	for ; y < yp+d; y++ {
		g.refreshSquare(x, y)
	}
	for ; x > xp-d; x-- {
		g.refreshSquare(x, y)
	}
	for ; y > yp-d; y-- {
		g.refreshSquare(x, y)
	}
}

func (g *game) refreshFrom(x, y int) {
	g.refreshSquare(x, y)

	for i := 1; i < fyne.Max(g.board.width, g.board.height); i++ {
		g.refreshAround(x, y, i)
	}
}

func (g *game) createRenderer() *gameRenderer {
	renderer := &gameRenderer{game: g}

	buttons := []fyne.CanvasObject{}
	for y := 0; y < g.board.height; y++ {
		for x := 0; x < g.board.width; x++ {
			xx, yy := x, y

			buttons = append(buttons, widget.NewButtonWithIcon(" ", code, func() {
				g.squareTapped(xx, yy)
			}))
		}
	}

	renderer.grid = fyne.NewContainerWithLayout(layout.NewGridLayout(g.board.width), buttons...)
	return renderer
}

func (g *game) squareTapped(x, y int) {
	g.board.reveal(x, y)
	g.refreshFrom(x, y)
}

func (g *game) loseCallback(yes bool) {
	if !yes {
		return
	}

	g.board.load(10)
	g.refreshFrom(g.board.width/2, g.board.height/2)
}

func (g *game) win() {
	dialog.ShowInformation("You won!", "Congratulations, you found all the bugs", g.window)
}

func (g *game) lose() {
	dialog.ShowConfirm("You lost!", "You hit a bug and lost the game, try again?", g.loseCallback, g.window)
}

func newGame(f *board) *game {
	g := &game{board: f}

	return g
}

// Show starts a new bugs game
func Show(app fyne.App) {
	b := newBoard(9, 9)
	game := newGame(b)

	b.win = game.win
	b.lose = game.lose
	b.load(10)

	game.window = app.NewWindow("Bugs")
	game.window.SetContent(game)

	game.window.Show()
}
