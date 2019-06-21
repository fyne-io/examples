// Package main launches the text editor example directly
package main

import (
	"fyne.io/fyne/app"

	"github.com/fyne-io/examples/img/icon"
	"github.com/fyne-io/examples/textedit"
)

func main() {
	app := app.New()
	app.SetIcon(icon.TextEditorBitmap)

	textedit.Show(app)
	app.Run()
}
