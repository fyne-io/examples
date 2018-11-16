// Package main launches the bugs example game directly
package main

import "github.com/fyne-io/examples/bugs"
import "github.com/fyne-io/fyne/app"

func main() {
	app := app.New()

	bugs.Show(app)
	app.Run()
}
