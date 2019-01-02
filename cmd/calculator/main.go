// Package main launches the calculator example directly
package main

import "fyne.io/fyne/app"

import "github.com/fyne-io/examples/calculator"

func main() {
	app := app.New()

	calculator.Show(app)
	app.Run()
}
