// Package main loads the sudoku example UI
package main

import (
	"fyne.io/fyne/app"
	"github.com/fyne-io/examples/img/icon"

	"github.com/fyne-io/examples/sudoku"
)

func main() {
	a := app.New()
	a.SetIcon(icon.SudokuBitmap)

	sudoku.Show(a)
	a.Run()
}
