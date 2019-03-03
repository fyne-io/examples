// Package main launches the clock example directly
package main

import (
	"fyne.io/fyne/app"
	"github.com/fyne-io/examples/clock"
	"github.com/fyne-io/examples/img/icon"
)

func main() {
	app := app.New()
	app.SetIcon(icon.ClockBitmap)

	clock.Show(app)
	app.Run()
}
