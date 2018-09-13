// Package main loads a very basic faces.Card demo application
package main

import "github.com/fyne-io/fyne/desktop"
import "github.com/fyne-io/examples/solitaire"

func main() {
	game := solitaire.NewGame()
	game.Deal()

	app := desktop.NewApp()
	w := app.NewWindow("Solitaire")
	w.SetContent(solitaire.NewTable(game))

	w.Show()
}
