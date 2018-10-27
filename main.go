package main

import (
	"flag"
	"fmt"
	"github.com/fyne-io/examples/bugs"
	"github.com/fyne-io/examples/calculator"
	"github.com/fyne-io/examples/fractal"
	"github.com/fyne-io/examples/life"
	"github.com/fyne-io/examples/solitaire"
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/desktop"
	"github.com/fyne-io/fyne/layout"
	"github.com/fyne-io/fyne/widget"
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
	w.ShowAndRun()
}

func main() {
	apps = make(map[string]func(fyne.App))
	apps["calculator"] = calculator.Show
	apps["fractal"] = fractal.Show
	apps["life"] = life.Show
	apps["bugs"] = bugs.Show
	apps["solitaire"] = solitaire.Show

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

	app := desktop.NewApp()
	launch(app)
}
