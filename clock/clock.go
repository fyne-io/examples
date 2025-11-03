package clock

import (
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type clockLayout struct {
	hour, minute, second     *canvas.Line
	pips                     [12]*canvas.Line
	hourDot, secondDot, face *canvas.Circle

	canvas fyne.CanvasObject
	stop   bool
}

func (c *clockLayout) rotate(hand *canvas.Line, middle fyne.Position, facePosition float64, offset, length float32) {
	rotation := math.Pi * 2 / 60 * facePosition
	x2 := length * float32(math.Sin(rotation))
	y2 := -length * float32(math.Cos(rotation))

	offX := float32(0)
	offY := float32(0)
	if offset > 0 {
		offX += offset * float32(math.Sin(rotation))
		offY += -offset * float32(math.Cos(rotation))
	}

	hand.Position1 = fyne.NewPos(middle.X+offX, middle.Y+offY)
	hand.Position2 = fyne.NewPos(middle.X+offX+x2, middle.Y+offY+y2)
	hand.Refresh()
}

func (c *clockLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	dotRadius := radius / 12
	smallDotRadius := dotRadius / 8

	stroke := diameter / 40
	midStroke := diameter / 90
	smallStroke := diameter / 200

	size = fyne.NewSize(diameter, diameter)
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	topleft := fyne.NewPos(middle.X-radius, middle.Y-radius)

	c.face.Resize(size)
	c.face.Move(topleft)

	c.hour.StrokeWidth = stroke
	c.rotate(c.hour, middle, float64((time.Now().Hour()%12)*5)+(float64(time.Now().Minute())/12), dotRadius, radius/2)
	c.minute.StrokeWidth = midStroke
	c.rotate(c.minute, middle, float64(time.Now().Minute())+(float64(time.Now().Second())/60), dotRadius, radius*.9)
	c.second.StrokeWidth = smallStroke
	c.rotate(c.second, middle, float64(time.Now().Second()), 0, radius-3)

	c.hourDot.StrokeWidth = stroke
	c.hourDot.Resize(fyne.NewSize(dotRadius*2, dotRadius*2))
	c.hourDot.Move(fyne.NewPos(middle.X-dotRadius, middle.Y-dotRadius))
	c.secondDot.StrokeWidth = smallStroke
	c.secondDot.Resize(fyne.NewSize(smallDotRadius*2, smallDotRadius*2))
	c.secondDot.Move(fyne.NewPos(middle.X-smallDotRadius, middle.Y-smallDotRadius))
	c.face.StrokeWidth = smallStroke

	for i, p := range c.pips {
		c.rotate(p, middle, float64((i)*5), radius/8*7, radius/8)
		p.StrokeWidth = smallStroke
	}
}

func (c *clockLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 200)
}

func (c *clockLayout) render() *fyne.Container {
	c.hourDot = &canvas.Circle{StrokeColor: theme.Color(theme.ColorNameForeground), StrokeWidth: 5}
	c.secondDot = &canvas.Circle{StrokeColor: theme.Color(theme.ColorNamePrimary), StrokeWidth: 3}
	c.face = &canvas.Circle{StrokeColor: theme.Color(theme.ColorNameDisabled), StrokeWidth: 1}

	c.hour = &canvas.Line{StrokeColor: theme.Color(theme.ColorNameForeground), StrokeWidth: 5}
	c.minute = &canvas.Line{StrokeColor: theme.Color(theme.ColorNameForeground), StrokeWidth: 3}
	c.second = &canvas.Line{StrokeColor: theme.Color(theme.ColorNamePrimary), StrokeWidth: 1}

	container := container.NewWithoutLayout(c.hourDot, c.secondDot)
	for i := range c.pips {
		pip := &canvas.Line{StrokeColor: theme.Color(theme.ColorNameDisabled), StrokeWidth: 1}
		container.Add(pip)
		c.pips[i] = pip
	}
	container.Objects = append(container.Objects, c.face, c.hour, c.minute, c.second)
	container.Layout = c

	c.canvas = container
	return container
}

func (c *clockLayout) animate(co fyne.CanvasObject) {
	tick := time.NewTicker(time.Second)
	go func() {
		for !c.stop {
			fyne.Do(func() {
				c.Layout(nil, co.Size())
			})
			<-tick.C
		}
	}()
}

func (c *clockLayout) applyTheme(_ fyne.Settings) {
	c.hourDot.StrokeColor = theme.Color(theme.ColorNameForeground)
	c.secondDot.StrokeColor = theme.Color(theme.ColorNamePrimary)
	c.face.StrokeColor = theme.Color(theme.ColorNameDisabled)

	c.hour.StrokeColor = theme.Color(theme.ColorNameForeground)
	c.minute.StrokeColor = theme.Color(theme.ColorNameForeground)
	c.second.StrokeColor = theme.Color(theme.ColorNamePrimary)

	for _, p := range c.pips {
		p.StrokeColor = theme.Color(theme.ColorNameDisabled)
	}
}

// Show loads a clock example window for the specified app context
func Show(win fyne.Window) fyne.CanvasObject {
	clock := &clockLayout{}
	win.SetOnClosed(func() {
		clock.stop = true
	})

	content := clock.render()
	go clock.animate(content)

	fyne.CurrentApp().Settings().AddListener(func(settings fyne.Settings) {
		clock.applyTheme(settings)
	})

	return content
}
