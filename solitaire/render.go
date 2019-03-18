package solitaire

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"

	"github.com/fyne-io/examples/solitaire/faces"
)

const minCardWidth = 95
const minPadding = 4
const cardRatio = 142.0 / minCardWidth

var (
	cardSize = fyne.Size{Width: minCardWidth, Height: minCardWidth * cardRatio}

	smallPad = minPadding
	overlap  = smallPad * 5
	bigPad   = smallPad + overlap

	minWidth  = cardSize.Width*7 + smallPad*8
	minHeight = cardSize.Height*3 + bigPad + smallPad*2
)

func updateCardPosition(c *canvas.Image, x, y int) {
	c.Resize(cardSize)
	c.Move(fyne.NewPos(x, y))
}

func withinCardBounds(c *canvas.Image, pos fyne.Position) bool {
	if pos.X < c.Position().X || pos.Y < c.Position().Y {
		return false
	}

	if pos.X >= c.Position().X+c.Size().Width || pos.Y >= c.Position().Y+c.Size().Height {
		return false
	}

	return true
}

func newCardPos(card *Card) *canvas.Image {
	if card == nil {
		return &canvas.Image{}
	}

	var face fyne.Resource
	if card.FaceUp {
		face = card.Face()
	} else {
		face = faces.ForBack()
	}
	image := &canvas.Image{Resource: face}
	image.Resize(cardSize)

	return image
}

func newCardSpace() *canvas.Image {
	space := faces.ForSpace()
	image := &canvas.Image{Resource: space}
	image.Resize(cardSize)

	return image
}

type tableRender struct {
	game *Game

	deck *canvas.Image

	pile1, pile2, pile3            *canvas.Image
	build1, build2, build3, build4 *canvas.Image

	stack1, stack2, stack3, stack4, stack5, stack6, stack7 *stackRender

	objects []fyne.CanvasObject
	table   *Table
}

func updateSizes(pad int) {
	smallPad = pad
	overlap = smallPad * 5
	bigPad = smallPad + overlap
}

func (t *tableRender) MinSize() fyne.Size {
	return fyne.NewSize(minWidth, minHeight)
}

func (t *tableRender) Layout(size fyne.Size) {
	padding := int(float32(size.Width) * .006)
	updateSizes(padding)

	newWidth := (size.Width - smallPad*8) / 7.0
	newWidth = fyne.Max(newWidth, minCardWidth)
	cardSize = fyne.NewSize(newWidth, int(float32(newWidth)*cardRatio))

	updateCardPosition(t.deck, smallPad, smallPad)

	updateCardPosition(t.pile1, smallPad*2+cardSize.Width, smallPad)
	updateCardPosition(t.pile2, smallPad*2+cardSize.Width+overlap, smallPad)
	updateCardPosition(t.pile3, smallPad*2+cardSize.Width+overlap*2, smallPad)

	updateCardPosition(t.build1, size.Width-(smallPad+cardSize.Width)*4, smallPad)
	updateCardPosition(t.build2, size.Width-(smallPad+cardSize.Width)*3, smallPad)
	updateCardPosition(t.build3, size.Width-(smallPad+cardSize.Width)*2, smallPad)
	updateCardPosition(t.build4, size.Width-(smallPad+cardSize.Width), smallPad)

	t.stack1.Layout(fyne.NewPos(smallPad, smallPad+bigPad+cardSize.Height),
		fyne.NewSize(cardSize.Width, size.Height-(smallPad+bigPad+cardSize.Height)))
	t.stack2.Layout(fyne.NewPos(smallPad+(smallPad+cardSize.Width), smallPad+bigPad+cardSize.Height),
		fyne.NewSize(cardSize.Width, size.Height-(smallPad+bigPad+cardSize.Height)))
	t.stack3.Layout(fyne.NewPos(smallPad+(smallPad+cardSize.Width)*2, smallPad+bigPad+cardSize.Height),
		fyne.NewSize(cardSize.Width, size.Height-(smallPad+bigPad+cardSize.Height)))
	t.stack4.Layout(fyne.NewPos(smallPad+(smallPad+cardSize.Width)*3, smallPad+bigPad+cardSize.Height),
		fyne.NewSize(cardSize.Width, size.Height-(smallPad+bigPad+cardSize.Height)))
	t.stack5.Layout(fyne.NewPos(smallPad+(smallPad+cardSize.Width)*4, smallPad+bigPad+cardSize.Height),
		fyne.NewSize(cardSize.Width, size.Height-(smallPad+bigPad+cardSize.Height)))
	t.stack6.Layout(fyne.NewPos(smallPad+(smallPad+cardSize.Width)*5, smallPad+bigPad+cardSize.Height),
		fyne.NewSize(cardSize.Width, size.Height-(smallPad+bigPad+cardSize.Height)))
	t.stack7.Layout(fyne.NewPos(smallPad+(smallPad+cardSize.Width)*6, smallPad+bigPad+cardSize.Height),
		fyne.NewSize(cardSize.Width, size.Height-(smallPad+bigPad+cardSize.Height)))
}

func (t *tableRender) ApplyTheme() {
	// no-op we are a custom UI
}

func (t *tableRender) BackgroundColor() color.Color {
	return color.RGBA{0x07, 0x63, 0x24, 0xff}
}

func (t *tableRender) refreshCard(img *canvas.Image, card *Card) {
	img.Hidden = card == nil
	t.refreshCardOrBlank(img, card)
}

func (t *tableRender) refreshCardOrBlank(img *canvas.Image, card *Card) {
	img.Resource = faces.ForSpace()
	img.Translucency = 0
	if card == nil {
		img.Resource = faces.ForSpace()
		return
	}

	if card != nil {
		if card.FaceUp {
			img.Resource = card.Face()
		} else {
			img.Resource = faces.ForBack()
		}

		if t.table.selected != nil && cardEquals(card, t.table.selected) {
			img.Translucency = 0.25
		} else {
			img.Translucency = 0
		}
	}
}

func (t *tableRender) Refresh() {
	if len(t.game.Hand.Cards) > 0 {
		t.deck.Resource = faces.ForBack()
	} else {
		t.deck.Resource = faces.ForSpace()
	}
	canvas.Refresh(t.deck)

	t.refreshCard(t.pile1, t.game.Draw1)
	t.refreshCard(t.pile2, t.game.Draw2)
	t.refreshCard(t.pile3, t.game.Draw3)

	t.refreshCardOrBlank(t.build1, t.game.Build1.Top())
	t.refreshCardOrBlank(t.build2, t.game.Build2.Top())
	t.refreshCardOrBlank(t.build3, t.game.Build3.Top())
	t.refreshCardOrBlank(t.build4, t.game.Build4.Top())

	t.stack1.Refresh(t.game.Stack1)
	t.stack2.Refresh(t.game.Stack2)
	t.stack3.Refresh(t.game.Stack3)
	t.stack4.Refresh(t.game.Stack4)
	t.stack5.Refresh(t.game.Stack5)
	t.stack6.Refresh(t.game.Stack6)
	t.stack7.Refresh(t.game.Stack7)

	canvas.Refresh(t.table)
}

func (t *tableRender) Objects() []fyne.CanvasObject {
	return t.objects
}

func (t *tableRender) Destroy() {
}

func (t *tableRender) appendStack(stack *stackRender) {
	for _, card := range stack.cards {
		t.objects = append(t.objects, card)
	}
}

func (t *tableRender) positionForCard(card *Card) *canvas.Image {

	return nil
}

func newTableRender(table *Table) *tableRender {
	render := &tableRender{}
	render.table = table
	render.game = table.game
	render.deck = newCardPos(nil)

	render.pile1 = newCardPos(nil)
	render.pile2 = newCardPos(nil)
	render.pile3 = newCardPos(nil)

	render.build1 = newCardSpace()
	render.build2 = newCardSpace()
	render.build3 = newCardSpace()
	render.build4 = newCardSpace()

	render.stack1 = newStackRender(render)
	render.stack2 = newStackRender(render)
	render.stack3 = newStackRender(render)
	render.stack4 = newStackRender(render)
	render.stack5 = newStackRender(render)
	render.stack6 = newStackRender(render)
	render.stack7 = newStackRender(render)

	render.objects = []fyne.CanvasObject{render.deck, render.pile1, render.pile2, render.pile3,
		render.build1, render.build2, render.build3, render.build4}

	render.appendStack(render.stack1)
	render.appendStack(render.stack2)
	render.appendStack(render.stack3)
	render.appendStack(render.stack4)
	render.appendStack(render.stack5)
	render.appendStack(render.stack6)
	render.appendStack(render.stack7)

	render.Refresh()
	return render
}

type stackRender struct {
	cards [19]*canvas.Image
	table *tableRender
}

func (s *stackRender) Layout(pos fyne.Position, size fyne.Size) {
	top := pos.Y
	for i := range s.cards {
		updateCardPosition(s.cards[i], pos.X, top)

		top += overlap
	}
}

func (s *stackRender) Refresh(stack *Stack) {
	var i int
	var card *Card
	if len(stack.Cards) == 0 {
		s.cards[0].Resource = faces.ForSpace()
		s.cards[0].Translucency = 0
		i = 0
	} else {
		for i, card = range stack.Cards {
			s.table.refreshCard(s.cards[i], card)
		}
	}

	for i = i + 1; i < len(s.cards); i++ {
		s.cards[i].Hide()
	}
}

func newStackRender(table *tableRender) *stackRender {
	r := &stackRender{table: table}
	for i := range r.cards {
		r.cards[i] = newCardPos(nil)
	}

	return r
}
