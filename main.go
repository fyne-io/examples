package main

import (
	"flag"
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
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
	run  func(fyne.App)
}

var apps []appInfo

func welcome(app fyne.App) {
	w := app.NewWindow("Examples")
	grid := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	group1 := widget.NewGroup("Canvas")
	grid.AddObject(group1)
	group2 := widget.NewGroup("Widgets")
	grid.AddObject(group2)

	for _, launch := range apps {
		list := group1
		if !launch.canv {
			list = group2
		}
		list.Append(newExampleButton(launch, app))
	}

	w.SetContent(grid)
	w.Show()
}

func flagFromName(name string) string {
	lower := strings.ToLower(name)
	return strings.Replace(lower, " ", "", -1)
}

func main() {
	apps = append(apps, appInfo{"Bugs", icon.BugBitmap, false, bugs.Show})
	apps = append(apps, appInfo{"Text Editor", icon.TextEditorBitmap, false, textedit.Show})
	apps = append(apps, appInfo{"XKCD", icon.XKCDBitmap, false, xkcd.Show})
	apps = append(apps, appInfo{"Clock", icon.ClockBitmap, true, clock.Show})
	apps = append(apps, appInfo{"Fractal", icon.FractalBitmap, true, fractal.Show})

	flags := make(map[string]*bool)
	for _, launch := range apps {
		name := flagFromName(launch.name)
		flags[name] = flag.Bool(name, false, fmt.Sprintf("Launch %s app directly", name))
	}
	flag.Parse()

	launch := welcome
	for ex, set := range flags {
		if *set {
			for _, a := range apps {
				if flagFromName(a.name) == ex {
					launch = a.run
				}
			}
			break
		}
	}

	ex := app.New()
	launch(ex)
	ex.Run()
}

type exampleButtonRenderer struct {
	icon  *canvas.Image
	label *canvas.Text

	objects []fyne.CanvasObject
	button  *exampleButton
}

func (b *exampleButtonRenderer) MinSize() fyne.Size {
	baseSize := b.label.MinSize()
	baseSize = baseSize.Add(fyne.NewSize(24, 24))
	return baseSize.Add(fyne.NewSize(theme.Padding()*4, theme.Padding()*2))
}

func (b *exampleButtonRenderer) Layout(size fyne.Size) {
	inner := size.Subtract(fyne.NewSize(theme.Padding()*4, theme.Padding()*2))
	inner = inner.Subtract(fyne.NewSize(24, 24))
	height := b.button.Size().Height

	b.label.Resize(inner)
	b.label.Move(fyne.NewPos(theme.Padding()*2+24, theme.Padding()+24))

	if b.icon != nil {
		b.icon.Resize(fyne.NewSize(height, height))
		b.icon.Move(fyne.NewPos(0, 0))
	}
}

func (b *exampleButtonRenderer) ApplyTheme() {
	b.label.Color = theme.TextColor()
	b.label.TextSize = theme.TextSize() * 2

	b.Refresh()
}

func (b *exampleButtonRenderer) BackgroundColor() color.Color {
	return theme.ButtonColor()
}

func (b *exampleButtonRenderer) Refresh() {
	b.Layout(b.button.Size())
	canvas.Refresh(b.button)
}

func (b *exampleButtonRenderer) Objects() []fyne.CanvasObject {
	return b.objects
}

func (b *exampleButtonRenderer) Destroy() {
}

type exampleButton struct {
	widget.BaseWidget
	Text string
	Icon fyne.Resource

	OnTapped func()
}

func (b *exampleButton) Tapped(*fyne.PointEvent) {
	b.OnTapped()
}

func (b *exampleButton) TappedSecondary(*fyne.PointEvent) {
}

func (b *exampleButton) CreateRenderer() fyne.WidgetRenderer {
	var objects []fyne.CanvasObject
	var img *canvas.Image
	if b.Icon != nil {
		img = canvas.NewImageFromResource(b.Icon)
		img.Translucency = 0.25
		objects = append(objects, img)
	}

	text := canvas.NewText(b.Text, theme.TextColor())
	text.TextSize = theme.TextSize() * 2
	text.Alignment = fyne.TextAlignTrailing

	objects = append(objects, text)
	return &exampleButtonRenderer{img, text, objects, b}
}

func newExampleButton(info appInfo, app fyne.App) *exampleButton {
	button := &exampleButton{Text: info.name, Icon: info.icon, OnTapped: func() {
		info.run(app)
	}}
	button.ExtendBaseWidget(button)

	return button
}
