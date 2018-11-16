// Package main launches a basic solitaire application
package main

import "github.com/fyne-io/examples/solitaire"
import "github.com/fyne-io/fyne/app"

func main() {
	app := app.New()

	solitaire.Show(app)
	app.Run()
}
