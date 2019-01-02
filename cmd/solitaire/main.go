// Package main launches a basic solitaire application
package main

import (
	"fyne.io/fyne/app"

	"github.com/fyne-io/examples/solitaire"
)

func main() {
	app := app.New()

	solitaire.Show(app)
	app.Run()
}
