package tictactoe

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Show loads a tic-tac-toe example window for the specified app context
func Show(win fyne.Window) fyne.CanvasObject {
	board := &board{}

	grid := container.NewGridWithColumns(3)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			grid.Add(newBoardIcon(r, c, board))
		}
	}

	reset := widget.NewButtonWithIcon("Reset Board", theme.ViewRefreshIcon(), func() {
		for i := range grid.Objects {
			grid.Objects[i].(*boardIcon).Reset()
		}

		board.Reset()
	})

	return container.NewBorder(reset, nil, nil, nil, grid)
}
