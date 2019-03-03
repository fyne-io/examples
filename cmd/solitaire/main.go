// Package main launches a basic solitaire application
package main

import (
	"fyne.io/fyne/app"
	"github.com/fyne-io/examples/img/icon"

	"github.com/fyne-io/examples/solitaire"
)

func main() {
	app := app.New()
	app.SetIcon(icon.SolitaireBitmap)

	solitaire.Show(app)
	app.Run()
}
