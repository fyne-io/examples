package fractal

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
)

type fractal struct {
	currIterations          uint
	currScale, currX, currY float64

	window fyne.Window
	canvas fyne.CanvasObject
}

func (f *fractal) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	f.canvas.Resize(size)
}

func (f *fractal) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(320, 240)
}

func (f *fractal) refresh() {
	if f.currScale >= 1.0 {
		f.currIterations = 100
	} else {
		f.currIterations = uint(100 * (1 + math.Pow((math.Log10(1/f.currScale)), 1.25)))
	}

	f.window.Canvas().Refresh(f.canvas)
}

func (f *fractal) scaleChannel(c float64, start, end uint32) uint8 {
	if end >= start {
		return (uint8)(c*float64(uint8(end-start))) + uint8(start)
	}

	return (uint8)((1-c)*float64(uint8(start-end))) + uint8(end)
}

func (f *fractal) scaleColor(c float64, start, end color.Color) color.Color {
	r1, g1, b1, _ := start.RGBA()
	r2, g2, b2, _ := end.RGBA()
	return color.RGBA{f.scaleChannel(c, r1, r2), f.scaleChannel(c, g1, g2), f.scaleChannel(c, b1, b2), 0xff}
}

func (f *fractal) mandelbrot(px, py, w, h int) color.Color {
	drawScale := 3.5 * f.currScale
	aspect := (float64(h) / float64(w))
	cRe := ((float64(px)/float64(w))-0.5)*drawScale + f.currX
	cIm := ((float64(py)/float64(w))-(0.5*aspect))*drawScale - f.currY

	var i uint
	var x, y, xsq, ysq float64

	for i = 0; i < f.currIterations && (xsq+ysq <= 4); i++ {
		xNew := float64(xsq-ysq) + cRe
		y = 2*x*y + cIm
		x = xNew

		xsq = x * x
		ysq = y * y
	}

	if i == f.currIterations {
		return theme.BackgroundColor()
	}

	mu := (float64(i) / float64(f.currIterations))
	c := math.Sin((mu / 2) * math.Pi)

	return f.scaleColor(c, theme.PrimaryColor(), theme.TextColor())
}

func (f *fractal) fractalRune(r rune) {
	if r == '+' {
		f.currScale /= 1.1
	} else if r == '-' {
		f.currScale *= 1.1
	}

	f.refresh()
}

func (f *fractal) fractalKey(ev *fyne.KeyEvent) {
	delta := f.currScale * 0.2
	if ev.Name == fyne.KeyUp {
		f.currY -= delta
	} else if ev.Name == fyne.KeyDown {
		f.currY += delta
	} else if ev.Name == fyne.KeyLeft {
		f.currX += delta
	} else if ev.Name == fyne.KeyRight {
		f.currX -= delta
	}

	f.refresh()
}

// Show loads a Mandelbrot fractal example window for the specified app context
func Show(win fyne.Window) fyne.CanvasObject {
	fractal := &fractal{window: win}
	fractal.canvas = canvas.NewRasterWithPixels(fractal.mandelbrot)

	fractal.currIterations = 100
	fractal.currScale = 1.0
	fractal.currX = -0.75
	fractal.currY = 0.0

	return fyne.NewContainerWithLayout(fractal, fractal.canvas)
	//TODO register, and unregister, these keys
	//window.Canvas().SetOnTypedRune(fractal.fractalRune)
	//window.Canvas().SetOnTypedKey(fractal.fractalKey)
}
