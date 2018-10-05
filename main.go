package main

import (
	"flag"
	"fmt"
	"github.com/fyne-io/fyne"

	"github.com/fyne-io/examples/calculator"
	"github.com/fyne-io/examples/fractal"
	"github.com/fyne-io/examples/life"
	"github.com/fyne-io/examples/solitaire"
	"github.com/fyne-io/fyne/desktop"
)

func welcome(_ fyne.App) {
	fmt.Println("Main UI not written, launch an app directly or use \"--help\"")
}

func main() {
	apps := make(map[string]func(fyne.App))
	apps["calculator"] = calculator.Show
	apps["fractal"] = fractal.Show
	apps["life"] = life.Show
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
