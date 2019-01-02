// Package main launches the fractal example directly
package main

import (
	"github.com/fyne-io/examples/fractal"
	"github.com/fyne-io/examples/img/icon"
)
import "fyne.io/fyne/app"

func main() {
	app := app.New()
	app.SetIcon(icon.FractalBitmap)

	fractal.Show(app)
	app.Run()
}
