//go:generate fyne bundle -o data.go Icon.png

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	"github.com/fyne-io/examples/bugs"
	"github.com/fyne-io/examples/clock"
	"github.com/fyne-io/examples/fractal"
	"github.com/fyne-io/examples/img/icon"
	"github.com/fyne-io/examples/textedit"
	"github.com/fyne-io/examples/xkcd"
)

type appInfo struct {
	name string
	icon fyne.Resource
	canv bool
	run  func(fyne.Window) fyne.CanvasObject
}

var apps = []appInfo{
	{"Bugs", icon.BugBitmap, false, bugs.Show},
	{"Text Editor", icon.TextEditorBitmap, false, textedit.Show},
	{"XKCD", icon.XKCDBitmap, false, xkcd.Show},
	{"Clock", icon.ClockBitmap, true, clock.Show},
	{"Fractal", icon.FractalBitmap, true, fractal.Show}}

func main() {
	a := app.New()
	a.SetIcon(resourceIconPng)

	content := container.NewMax()
	w := a.NewWindow("Examples")

	appList := widget.NewList(
		func() int {
			return len(apps)
		},
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, &canvas.Image{}, nil,
				widget.NewLabel("Text Editor"))
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[1].(*canvas.Image)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			img.Resource = apps[id].icon
			img.Refresh()
			text.SetText(apps[id].name)
		})
	appList.OnSelected = func(id widget.ListItemID) {
		content.Objects = []fyne.CanvasObject{apps[id].run(w)}
	}

	split := container.NewHSplit(appList, content)
	split.Offset = 0.1
	w.SetContent(split)
	w.Resize(fyne.NewSize(480, 360))
	w.ShowAndRun()
}
