package main

import (
	"flag"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/fyne-io/examples/bugs"
	"github.com/fyne-io/examples/calculator"
	"github.com/fyne-io/examples/fractal"
	"github.com/fyne-io/examples/life"
	"github.com/fyne-io/examples/solitaire"
	"github.com/fyne-io/examples/sudoku"
	"github.com/fyne-io/examples/xkcd"
)

var apps map[string]func(fyne.App)

func welcome(app fyne.App) {
	w := app.NewWindow("Examples")
	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(2))

	for name := range apps {
		launch := apps[name]
		grid.AddObject(widget.NewButton(name, func() {
			launch(app)
		}))
	}

	w.SetContent(grid)
	w.Show()
}

func main() {
	apps = make(map[string]func(fyne.App))
	apps["calculator"] = calculator.Show
	apps["fractal"] = fractal.Show
	apps["life"] = life.Show
	apps["bugs"] = bugs.Show
	apps["solitaire"] = solitaire.Show
	apps["xkcd"] = xkcd.Show
	apps["sudoku"] = sudoku.Show

	flags := make(map[string]*bool)
	for name := range apps {
		flags[name] = flag.Bool(name, false, fmt.Sprintf("Launch %s app directly", name))
	}
	flag.Parse()

	launch := welcome
	for app, set := range flags {
		if *set {
			launch = apps[app]
			break
		}
	}

	app := app.New()
	launch(app)
	app.Run()
}
