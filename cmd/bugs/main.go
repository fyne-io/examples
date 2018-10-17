// Package main launches the bugs example game directly
package main

import "github.com/fyne-io/examples/bugs"
import "github.com/fyne-io/fyne/desktop"

func main() {
	app := desktop.NewApp()

	bugs.Show(app)
}
