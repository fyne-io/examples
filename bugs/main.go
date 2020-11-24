package bugs

import (
	"fmt"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var bug, code, flag *theme.ThemedResource

func init() {
	bug = theme.NewThemedResource(bugIcon, nil)
	code = theme.NewThemedResource(codeIcon, nil)
	flag = theme.NewThemedResource(flagIcon, nil)
}

type gameRenderer struct {
	grid   *fyne.Container
	header fyne.CanvasObject

	game *game
}

func (g *gameRenderer) MinSize() fyne.Size {
	return g.grid.MinSize().Add(fyne.NewSize(0, g.header.MinSize().Height))
}

func (g *gameRenderer) Layout(size fyne.Size) {
	headerHeight := g.header.MinSize().Height
	g.header.Resize(fyne.NewSize(size.Width, headerHeight))
	g.grid.Move(fyne.NewPos(0, headerHeight)) // TODO why ignored?
	gridSize := size.Subtract(fyne.NewSize(0, headerHeight))
	g.grid.Layout.Layout(g.grid.Objects, gridSize)
}

func (g *gameRenderer) ApplyTheme() {
}

func (g *gameRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (g *gameRenderer) Refresh() {
	canvas.Refresh(g.grid)
}

func (g *gameRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{g.grid, g.header}
}

func (g *gameRenderer) Destroy() {
}

type game struct {
	widget.BaseWidget
	board  *board
	remain *widget.Label

	window fyne.Window
}

func (g *game) refreshSquare(x, y int) {
	if x < 0 || y < 0 || x >= g.board.width || y >= g.board.height {
		return
	}

	sq := g.board.bugs[y][x]
	i := y*g.board.width + x
	button := widget.Renderer(g).(*gameRenderer).grid.Objects[i].(*bugButton)

	if sq.flagged {
		if button.icon == flag {
			return
		}
		button.icon = flag
		button.text = ""
	} else if !sq.shown {
		if button.icon == code {
			return
		}
		button.icon = code
		button.text = ""
	} else if sq.bug {
		if button.icon == bug {
			return
		}
		button.icon = bug
		button.text = ""
	} else if button.icon == nil {
		return
	} else {
		button.icon = nil
		button.text = squareString(sq)
	}

	widget.Refresh(button)
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

func (g *game) CreateRenderer() fyne.WidgetRenderer {
	renderer := &gameRenderer{game: g}
	title := widget.NewLabel("Hunt bugs!")
	g.remain = widget.NewLabel("")
	g.updateRemain()
	renderer.header = fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, title, g.remain),
		title, g.remain)

	var buttons []fyne.CanvasObject
	for y := 0; y < g.board.height; y++ {
		for x := 0; x < g.board.width; x++ {
			xx, yy := x, y

			buttons = append(buttons, newButton("", code, func(reveal bool) {
				if reveal {
					g.squareReveal(xx, yy)
				} else {
					g.squareFlagged(xx, yy)
				}

				g.updateRemain()
			}))
		}
	}

	renderer.grid = fyne.NewContainerWithLayout(layout.NewGridLayout(g.board.width), buttons...)
	return renderer
}

func (g *game) squareReveal(x, y int) {
	if g.board.flagged(x, y) {
		return
	}

	g.board.reveal(x, y)
	g.refreshFrom(x, y)
}

func (g *game) squareFlagged(x, y int) {
	g.board.flag(x, y)
	g.refreshSquare(x, y)
}

func (g *game) loseCallback(yes bool) {
	if !yes {
		return
	}

	g.board.load(40)
	g.updateRemain()
	g.refreshFrom(g.board.width/2, g.board.height/2)
}

func (g *game) win() {
	dialog.ShowInformation("You won!", "Congratulations, you found all the bugs", g.window)
}

func (g *game) lose() {
	dialog.ShowConfirm("You lost!", "You hit a bug and lost the game, try again?", g.loseCallback, g.window)
}

func (g *game) updateRemain() {
	g.remain.SetText(fmt.Sprintf("remaining: %d", g.board.remaining()))
}

func newGame(f *board) *game {
	g := &game{board: f}
	g.ExtendBaseWidget(g)

	return g
}

// Show starts a new bugs game
func Show(win fyne.Window) fyne.CanvasObject {
	b := newBoard(20, 14)
	game := newGame(b)

	b.win = game.win
	b.lose = game.lose
	b.load(40)

	game.window = win
	return game
}
