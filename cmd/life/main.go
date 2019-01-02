// Package main launches the game of life example directly
package main

import (
	"github.com/fyne-io/examples/img/icon"
	"github.com/fyne-io/examples/life"
)
import "fyne.io/fyne/app"

func main() {
	app := app.New()
	app.SetIcon(icon.LifeBitmap)

	life.Show(app)
	app.Run()
}
