// Package main launches the fractal example directly
package main

import "github.com/fyne-io/examples/fractal"
import "github.com/fyne-io/fyne/app"

func main() {
	app := app.New()

	fractal.Show(app)
	app.Run()
}
