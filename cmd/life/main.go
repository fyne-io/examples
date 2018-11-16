// Package main launches the game of life example directly
package main

import "github.com/fyne-io/examples/life"
import "github.com/fyne-io/fyne/app"

func main() {
	app := app.New()

	life.Show(app)
	app.Run()
}
