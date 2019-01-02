// Package main launches the bugs example game directly
package main

import (
	"fyne.io/fyne/app"

	"github.com/fyne-io/examples/bugs"
	"github.com/fyne-io/examples/img/icon"
)

func main() {
	app := app.New()
	app.SetIcon(icon.BugBitmap)

	bugs.Show(app)
	app.Run()
}
