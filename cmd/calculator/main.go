// Package main launches the calculator example directly
package main

import "github.com/fyne-io/examples/calculator"
import "github.com/fyne-io/fyne/app"

func main() {
	app := app.New()

	calculator.Show(app)
	app.Run()
}
