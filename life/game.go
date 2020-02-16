package life

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

var (
	pixDensity = 1.0
)

const (
	cellSize = 8
)

type gameRenderer struct {
	render   *canvas.Raster
	objects  []fyne.CanvasObject
	imgCache *image.RGBA

	aliveColor color.Color
	deadColor  color.Color

	game *game
}

func (g *gameRenderer) MinSize() fyne.Size {
	return fyne.NewSize(int(float64(minXCount*cellSize)/pixDensity), int(float64(minYCount*cellSize)/pixDensity))
}

func (g *gameRenderer) Layout(size fyne.Size) {
	g.render.Resize(size)
}

func (g *gameRenderer) ApplyTheme() {
	g.aliveColor = theme.TextColor()
	g.deadColor = theme.BackgroundColor()
}

func (g *gameRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (g *gameRenderer) Refresh() {
	canvas.Refresh(g.render)
}

func (g *gameRenderer) Objects() []fyne.CanvasObject {
	return g.objects
}

func (g *gameRenderer) Destroy() {
}

func (g *gameRenderer) draw(w, h int) image.Image {
	img := g.imgCache
	if img == nil || img.Bounds().Size().X != w || img.Bounds().Size().Y != h {
		img = image.NewRGBA(image.Rect(0, 0, w, h))
		g.imgCache = img
	}
	g.game.board.ensureGridSize(g.game.cellForCoord(w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xpos, ypos := g.game.cellForCoord(x, y)

			if xpos < g.game.board.width && ypos < g.game.board.height && g.game.board.cells[ypos][xpos] {
				img.Set(x, y, g.aliveColor)
			} else {
				img.Set(x, y, g.deadColor)
			}
		}
	}

	return img
}

type game struct {
	widget.BaseWidget

	genText *widget.Label
	board   *board
	paused  bool
}

func (g *game) CreateRenderer() fyne.WidgetRenderer {
	renderer := &gameRenderer{game: g}

	render := canvas.NewRaster(renderer.draw)
	renderer.render = render
	renderer.objects = []fyne.CanvasObject{render}
	renderer.ApplyTheme()

	return renderer
}

func (g *game) cellForCoord(x, y int) (int, int) {
	xpos := int(float64(x) / float64(cellSize) / pixDensity)
	ypos := int(float64(y) / float64(cellSize) / pixDensity)

	return xpos, ypos
}

func (g *game) run() {
	g.paused = false
}

func (g *game) stop() {
	g.paused = true
}

func (g *game) toggleRun() {
	g.paused = !g.paused
}

func (g *game) animate() {
	go func() {
		tick := time.NewTicker(time.Second / 6)

		for {
			select {
			case <-tick.C:
				if g.paused {
					continue
				}

				g.board.nextGen()
				g.updateGeneration()
				widget.Refresh(g)
			}
		}
	}()
}

func (g *game) typedRune(r rune) {
	if r == ' ' {
		g.toggleRun()
	}
}

func (g *game) Tapped(ev *fyne.PointEvent) {
	xpos, ypos := g.cellForCoord(int(float64(ev.Position.X)*pixDensity), int(float64(ev.Position.Y)*pixDensity))

	if ev.Position.X < 0 || ev.Position.Y < 0 || xpos >= g.board.width || ypos >= g.board.height {
		return
	}

	g.board.cells[ypos][xpos] = !g.board.cells[ypos][xpos]

	widget.Refresh(g)
}

func (g *game) TappedSecondary(ev *fyne.PointEvent) {
}

func (g *game) buildUI() fyne.CanvasObject {
	var pause *widget.Button
	pause = widget.NewButton("Pause", func() {
		g.paused = !g.paused

		if g.paused {
			pause.SetText("Resume")
		} else {
			pause.SetText("Pause")
		}
	})

	title := fyne.NewContainerWithLayout(layout.NewGridLayout(2), g.genText, pause)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(title, nil, nil, nil), title, g)
}

func (g *game) updateGeneration() {
	g.genText.SetText(fmt.Sprintf("Generation %d", g.board.generation))
}

func newGame(b *board) *game {
	g := &game{board: b, genText: widget.NewLabel("Generation 0")}
	g.ExtendBaseWidget(g)

	return g
}

func (g *game) adaptToTextureSize(c fyne.Canvas) {
	pixW, _ := c.PixelCoordinateForPosition(fyne.NewPos(cellSize, cellSize))
	pixDensity = float64(pixW) / float64(cellSize)
}
