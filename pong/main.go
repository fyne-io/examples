package pong

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

const (
	LEFT  = iota
	RIGHT = iota
	UP    = iota
	DOWN  = iota
)

type coord struct {
	x float64
	y float64
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
	b.location.x = float64(w / 2)
	b.location.y = float64(h / 2)
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
	dots        []bool
	w, h        int
	booting     bool
	glitchStart int
	glitchEnd   int
}

func newTV(w, h int) *tv {
	t := &tv{
		dots: make([]bool, w*h),
		w:    w,
		h:    h,
	}

	t.reboot()
	return t
}

func (t *tv) reboot() {
	t.booting = true

	// Init the static
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < (t.w * t.h); i++ {
		if rand.Intn(2) == 1 {
			t.dots[i] = true
		}
	}

	// the PotatoTV needs some random time to boot, its not very quick
	go func(t *tv) {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+2500))
		t.booting = false
	}(t)
}

func (t *tv) glitch(size fyne.Size) {
	if rand.Intn(100) == 1 {
		t.glitchStart = rand.Intn(size.Width * size.Height)
		t.glitchEnd = rand.Intn(size.Width * size.Height)
		if t.glitchStart > t.glitchEnd {
			t.glitchStart = 1
		}
		return
	}
	t.glitchStart = 0
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

	// does left score ?
	if g.ball.location.x < 1 {
		//memdebug.Print(time.Now(), "left !!")
		//time.Sleep(time.Second)
		//g.ball.newServe(g.tv.w, g.tv.h)
		g.ball.location.x = 1
		g.ball.location.y = float64(g.batLeft + (g.batHeight / 2))
		g.ball.velocity.x = 1
		g.ball.velocity.y = g.ball.velocity.y * 1.1
		return
	}

	// does right score ?
	if g.ball.location.x >= float64(g.tv.w) {
		//memdebug.Print(time.Now(), "right !!")
		//time.Sleep(time.Second)
		//g.ball.newServe(g.tv.w, g.tv.h)
		g.ball.location.x = float64(g.tv.w - 1)
		g.ball.location.y = float64(g.batRight + (g.batHeight / 2))
		g.ball.velocity.x = -1
		g.ball.velocity.y = g.ball.velocity.y * 1.1
		return
	}

	// top bounce ?
	if g.ball.location.y <= 8 {
		g.ball.location.y++
		g.ball.velocity.y *= -0.9
	}

	// bottom bounce ?
	if g.ball.location.y >= float64(g.tv.h-1) {
		g.ball.location.y--
		g.ball.velocity.y *= -0.9
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
	size      fyne.Size
}

func (g *gameRenderer) MinSize() fyne.Size {
	return fyne.NewSize(400, 300)
}

func (g *gameRenderer) Layout(size fyne.Size) {
	g.render.Resize(size)
	g.size = size
}

func (g *gameRenderer) ApplyTheme() {
	g.color = theme.TextColor()
	g.backColor = theme.BackgroundColor()
}

func (g *gameRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (g *gameRenderer) Refresh() {
	g.game.tv.glitch(g.size)
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

		// now rotate that static
		g.game.tv.dots = append(g.game.tv.dots[1:], rand.Intn(2) == 1)
		return retval
	}

	// static burst
	if g.game.tv.glitchStart > 0 {
		offset := y*w + x
		if offset >= g.game.tv.glitchStart && offset <= g.game.tv.glitchEnd {
			if rand.Intn(2) == 1 {
				return g.color
			}
			return g.backColor
		}
	}

	// so assume each pixel is black, and then look for reasons for the pixel to be lit

	// top 8 pixels is the scoreline
	// TODO - render scores here, using old school pixelmap

	// top bar at the 8th pixel
	if cy == 7 {
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
	if cx == int(g.game.ball.location.x) && cy == int(g.game.ball.location.y) {
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
	//memdebug.Print(time.Now(), "typed", *k)
	switch k.Name {
	case "Escape":
		g.tv.reboot()
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
	// TODO - we can do something in here maybe
}

// Show starts a new game of pong
func Show(app fyne.App) {
	game := newGame(64, 64)

	window := app.NewWindow("Potato Arcade Pong")
	window.SetContent(game)
	window.Canvas().SetOnTypedRune(game.typedRune)
	window.Canvas().SetOnTypedKey(game.typedKey)
	//window.Canvas().SetOnReleaseKey(game.releaseKey)

	// start the board animation before we show the window - it will block
	game.animate()
	window.Show()
}
