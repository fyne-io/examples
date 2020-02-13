package life

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/fyne-io/examples/img/icon"
)

const (
	minXCount = 50
	minYCount = 40
)

// Show starts a new game of life
func Show(app fyne.App) {
	board := newBoard(minXCount, minYCount)
	board.load()

	game := newGame(board)

	window := app.NewWindow("Life")
	window.SetIcon(icon.LifeBitmap)
	pause := widget.NewButton("Pause", func() {
		game.paused = !game.paused
	})
	window.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, pause, nil, nil), pause, game))
	window.Canvas().SetOnTypedRune(game.typedRune)
	game.adaptToTextureSize(window.Canvas())

	// start the board animation before we show the window - it will block
	game.animate()

	window.Show()
}
