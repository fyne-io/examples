// Package main launches the fractal example directly
package main

import "github.com/fyne-io/examples/fractal"
import "github.com/fyne-io/fyne/desktop"

func main() {
	app := desktop.NewApp()

	fractal.Show(app)
}
