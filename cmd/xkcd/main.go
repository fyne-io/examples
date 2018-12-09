// Package main launches the xkcd example game directly
package main

import "github.com/fyne-io/examples/xkcd"
import "github.com/fyne-io/fyne/app"

func main() {
	app := app.New()

	xkcd.Show(app)
	app.Run()
}
