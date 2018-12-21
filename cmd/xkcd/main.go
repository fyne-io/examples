// Package main launches the xkcd example game directly
package main

import (
	"github.com/fyne-io/fyne/app"

	"github.com/fyne-io/examples/img/icon"
	"github.com/fyne-io/examples/xkcd"
)

func main() {
	app := app.New()
	app.SetIcon(icon.XKCDBitmap)

	xkcd.Show(app)
	app.Run()
}
