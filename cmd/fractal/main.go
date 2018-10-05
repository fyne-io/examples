// Package main launches the fractal example directly
package main

import "github.com/fyne-io/examples/fractal"
import "github.com/fyne-io/fyne/examples/apps"

func main() {
	app := apps.NewApp()

	fractal.Show(app)
}
