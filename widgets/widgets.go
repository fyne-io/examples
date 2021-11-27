package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//
// Display both widgits and some helpfull text.
//
func Show(win fyne.Window) fyne.CanvasObject {
	// The Widget is 75 wide 75 high. Clock face is inset 4
	// Create the clock widget with a clock face from img/clock2.svg.
	// The clock is started.
	wc := NewWidgetClock(75, 75, 4, canvas.NewImageFromFile("img/clock2.svg"), true)

	// Create the Animated widget.
	// The widget is 75 wide 75 high.
	// When the widget is tapped the clock is stopped or started.
	wa := NewWidgetAnimate(75, 75, func(pe *fyne.PointEvent) {
		wc.SetRunning(!wc.GetRunning())
	})
	// Frame the widgets for the demo app.
	title1 := widget.NewLabel("Widgets created from scratch")
	title2 := widget.NewLabel("Widget with Animation")
	aw := container.NewHBox(wa, widget.NewLabel("Click to Animate"))
	title3 := widget.NewLabel("Simple clock widget")
	cw := container.NewHBox(wc, widget.NewLabel("Click circle above to STOP/START"))
	c := container.NewVBox(title1, title2, aw, title3, cw)
	return c
}
