// Package main launches the game of life example directly
package main

import "github.com/fyne-io/examples/life"
import "github.com/fyne-io/fyne/desktop"

func main() {
	app := desktop.NewApp()

	life.Show(app)
	app.Run()
}
