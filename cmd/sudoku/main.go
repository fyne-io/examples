// Package main loads the sudoku example UI
package main

import (
	"github.com/fyne-io/examples/sudoku"
	"github.com/fyne-io/fyne/app"
)

func main() {
	a := app.New()
	sudoku.Show(a)
	a.Run()
}
