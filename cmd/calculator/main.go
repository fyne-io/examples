// Package main launches the calculator example directly
package main

import (
	"fyne.io/fyne/app"
	"github.com/fyne-io/examples/calculator"
	"github.com/fyne-io/examples/img/icon"
)

func main() {
	app := app.New()
	app.SetIcon(icon.CalculatorBitmap)

	calculator.Show(app)
	app.Run()
}
