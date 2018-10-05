// Package main launches a basic solitaire application
package main

import "github.com/fyne-io/examples/solitaire"
import "github.com/fyne-io/fyne/desktop"

func main() {
	app := desktop.NewApp()

	solitaire.Show(app)
}
