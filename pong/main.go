package pong

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/steveoc64/memdebug"
)

const (
	LEFT  = iota
	RIGHT = iota
	UP    = iota
	DOWN  = iota
)

type coord struct {
	x int
	y int
}

type ball struct {
	location coord
	velocity coord
}

func (b *ball) move() {
	b.location.x += b.velocity.x
	b.location.y += b.velocity.y
}

func (b *ball) newServe(w, h int) {
	b.location.x = w / 2
	b.location.y = h / 2
	switch rand.Intn(4) {
	case 0:
		b.velocity = coord{x: 1, y: 1}
	case 1:
		b.velocity = coord{x: 1, y: -1}
	case 2:
		b.velocity = coord{x: -1, y: 1}
	case 3:
		b.velocity = coord{x: -1, y: -1}
	}
}

func newBall(w, h int) *ball {
	b := &ball{}
	b.newServe(w, h)
	return b
}

// tv emulation of a potato TV with a resolution of 64x64
type tv struct {
	dots    []bool
	w, h    int
	booting bool
}

func newTV(w, h int) *tv {
	t := &tv{
		dots: make([]bool, w*h),
		w:    w,
		h:    h,
	}
	// the PotatoTV needs some random time to boot
	go func(t *tv) {
		t.booting = true
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < (w * h); i++ {
			if rand.Intn(1) == 1 {
				t.dots[i] = true
			}
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+2500))
		t.booting = false
	}(t)
	return t
}

type game struct {
	tv        *tv
	paused    bool
	size      fyne.Size
	position  fyne.Position
	hidden    bool
	ball      *ball
	batLeft   int
	batRight  int
	batHeight int
}

func newGame(w, h int) *game {
	g := &game{
		tv:        newTV(w, h),
		ball:      newBall(w, h),
		batLeft:   h/2 - 4,
		batRight:  h/2 - 4,
		batHeight: 8,
	}
	return g
}

func (g *game) cellForCoord(x, y, w, h int) (int, int) {
	xpos := int(float64(g.tv.w) * (float64(x) / float64(w)))
	ypos := int(float64(g.tv.h) * (float64(y) / float64(h)))

	return xpos, ypos
}

func (g *game) moveBat(bat int, direction int) {
	switch bat {
	case LEFT:
		switch direction {
		case UP:
			if g.batLeft > 8 {
				g.batLeft--
			}
		case DOWN:
			if g.batLeft+g.batHeight < (g.tv.h - 1) {
				g.batLeft++
			}
		}
	case RIGHT:
		switch direction {
		case UP:
			if g.batRight > 8 {
				g.batRight--
			}
		case DOWN:
			if g.batRight+g.batHeight < (g.tv.h - 1) {
				g.batRight++
			}
		}
	}
}

func (g *game) moveBall() {
	g.ball.move()
	if g.ball.location.x < 1 {
		memdebug.Print(time.Now(), "left !!")
		time.Sleep(time.Second)
		g.ball.newServe(g.tv.w, g.tv.h)
	}
	if g.ball.location.x >= g.tv.w {
		memdebug.Print(time.Now(), "right !!")
		time.Sleep(time.Second)
		g.ball.newServe(g.tv.w, g.tv.h)

	}
}

////////////////////////////
// Boilerplate

func (g *game) Size() fyne.Size {
	return g.size
}

func (g *game) Resize(size fyne.Size) {
	g.size = size
	widget.Renderer(g).Layout(size)
}

func (g *game) Position() fyne.Position {
	return g.position
}

func (g *game) Move(pos fyne.Position) {
	g.position = pos
	widget.Renderer(g).Layout(g.size)
}

func (g *game) MinSize() fyne.Size {
	return widget.Renderer(g).MinSize()
}

func (g *game) Visible() bool {
	return g.hidden
}

func (g *game) Show() {
	g.hidden = false
}

func (g *game) Hide() {
	g.hidden = true
}

//////////////////////////////////////////
// rendering stuff

type gameRenderer struct {
	render    *canvas.Raster
	objects   []fyne.CanvasObject
	color     color.Color
	backColor color.Color
	game      *game
}

func (g *gameRenderer) MinSize() fyne.Size {
	return fyne.NewSize(400, 300)
}

func (g *gameRenderer) Layout(size fyne.Size) {
	g.render.Resize(size)
}

func (g *gameRenderer) ApplyTheme() {
	g.color = theme.TextColor()
	g.backColor = theme.BackgroundColor()
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

func (g *gameRenderer) renderer(x, y, w, h int) color.Color {
	cx, cy := g.game.cellForCoord(x, y, w, h)

	if g.game.tv.booting {
		// TV is booting, show static
		retval := g.backColor
		if g.game.tv.dots[cy*g.game.tv.w+cx] {
			retval = g.color
		}
		return retval
	}

	// TV is booted

	// so assume its black, and then look for reasons for the pixel to be lit

	// top 8 pixels is the scoreline

	// top bar at pixel 8
	if cy == 8 {
		return g.color
	}

	// lower bar at last pixel
	if cy == g.game.tv.h-1 {
		return g.color
	}

	// net - dotted line down the middle
	if cy > 8 && cx == g.game.tv.w/2 {
		if (cy % 2) == 0 {
			return g.color
		}
	}

	// left bat
	if cx == 0 && cy >= (g.game.batLeft) && cy <= (g.game.batLeft+g.game.batHeight) {
		return g.color
	}

	// right bat
	if cx == (g.game.tv.w-1) && cy >= (g.game.batRight) && cy <= (g.game.batRight+g.game.batHeight) {
		return g.color
	}

	// is it the ball ?
	if cx == g.game.ball.location.x && cy == g.game.ball.location.y {
		return g.color
	}

	return g.backColor
}

func (g *game) CreateRenderer() fyne.WidgetRenderer {
	renderer := &gameRenderer{game: g}

	render := canvas.NewRasterWithPixels(renderer.renderer)
	renderer.render = render
	renderer.objects = []fyne.CanvasObject{render}
	renderer.ApplyTheme()

	return renderer
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
		tick := time.NewTicker(time.Second / 12)

		for {
			select {
			case <-tick.C:
				if g.paused {
					continue
				}
				g.moveBall()
				widget.Refresh(g)
			}
		}
	}()
}

func (g *game) typedRune(r rune) {
	// vi motion keys - both qwerty && dvorak !!
	switch r {
	case ' ':
		g.toggleRun()
	case 'd', 'e':
		g.moveBat(LEFT, DOWN)
	case 'f', 'u':
		g.moveBat(LEFT, UP)
	case 'j', 'h':
		g.moveBat(RIGHT, DOWN)
	case 'k', 't':
		g.moveBat(RIGHT, UP)
	}
}

func (g *game) typedKey(k *fyne.KeyEvent) {
	memdebug.Print(time.Now(), "typed", *k)
	switch k.Name {
	case "Up":
		g.moveBat(RIGHT, UP)
	case "Down":
		g.moveBat(RIGHT, DOWN)
	case "Shift":
		g.moveBat(LEFT, UP)
	case "Control":
		g.moveBat(LEFT, DOWN)
	}
}

func (g *game) releaseKey(k *fyne.KeyEvent) {
	memdebug.Print(time.Now(), "released", *k)
}

// Show starts a new game of pong
func Show(app fyne.App) {
	game := newGame(64, 64)

	window := app.NewWindow("Pong")
	window.SetContent(game)
	window.Canvas().SetOnTypedRune(game.typedRune)
	window.Canvas().SetOnTypedKey(game.typedKey)
	window.Canvas().SetOnReleaseKey(game.releaseKey)

	// start the board animation before we show the window - it will block
	game.animate()
	window.Show()
}
