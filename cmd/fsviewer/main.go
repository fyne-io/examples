// Package main launches the fsviewer example directly
package main

import (
	"fyne.io/fyne/app"
	"github.com/fyne-io/examples/fsviewer"
	// TODO "github.com/fyne-io/examples/img/icon"
)

func main() {
	app := app.New()
	// TODO app.SetIcon(icon.FileViewerBitmap)

	fsviewer.Show(app)
	app.Run()
}
