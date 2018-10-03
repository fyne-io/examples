package solitaire

import (
	"github.com/fyne-io/examples/solitaire/faces"
	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/canvas"
	"image/color"
)

var cardSize = fyne.Size{Width: 95, Height: 142}

const smallPad = 4
const bigPad = 24
const overlap = 20

func newCard(face fyne.Resource) *canvas.Image {
	card := &canvas.Image{}
	if face != nil {
		card.File = face.CachePath()
	}
	card.Resize(cardSize)

	return card
}

type tableRender struct {
	game *Game

	bg   *canvas.Rectangle
	deck *canvas.Image

	pile1, pile2, pile3            *canvas.Image
	space1, space2, space3, space4 *canvas.Image

	stack1, stack2, stack3, stack4, stack5, stack6, stack7 *stackRender

	objects []fyne.CanvasObject
	table   *table
}

func (t *tableRender) MinSize() fyne.Size {
	return fyne.NewSize(cardSize.Width*7+smallPad*8, cardSize.Height*3+bigPad+smallPad*2)
}

func (t *tableRender) Layout(size fyne.Size) {
	t.bg.Resize(size)

	t.deck.Move(fyne.NewPos(smallPad, smallPad))

	t.pile1.Move(fyne.NewPos(smallPad*2+cardSize.Width, smallPad))
	t.pile2.Move(fyne.NewPos(smallPad*2+cardSize.Width+overlap, smallPad))
	t.pile3.Move(fyne.NewPos(smallPad*2+cardSize.Width+overlap*2, smallPad))

	t.space1.Move(fyne.NewPos(size.Width-(smallPad+cardSize.Width)*4, smallPad))
	t.space2.Move(fyne.NewPos(size.Width-(smallPad+cardSize.Width)*3, smallPad))
	t.space3.Move(fyne.NewPos(size.Width-(smallPad+cardSize.Width)*2, smallPad))
	t.space4.Move(fyne.NewPos(size.Width-(smallPad+cardSize.Width), smallPad))

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

func (t *tableRender) Refresh() {
	if len(t.game.Deck.Cards) > 0 {
		t.deck.File = faces.Back.CachePath()
	} else {
		t.deck.File = faces.Space.CachePath()
	}
	fyne.RefreshObject(t.deck)

	if t.game.Draw1 != nil {
		t.pile1.File = t.game.Draw1.Face().CachePath()
	}
	if t.game.Draw2 != nil {
		t.pile2.File = t.game.Draw2.Face().CachePath()
	}
	if t.game.Draw3 != nil {
		t.pile3.File = t.game.Draw3.Face().CachePath()
	}
	fyne.RefreshObject(t.pile1)
	fyne.RefreshObject(t.pile2)
	fyne.RefreshObject(t.pile3)

	t.stack1.Refresh(t.game.Stack1)
	t.stack2.Refresh(t.game.Stack2)
	t.stack3.Refresh(t.game.Stack3)
	t.stack4.Refresh(t.game.Stack4)
	t.stack5.Refresh(t.game.Stack5)
	t.stack6.Refresh(t.game.Stack6)
	t.stack7.Refresh(t.game.Stack7)

	fyne.RefreshObject(t.table)
}

func (t *tableRender) Objects() []fyne.CanvasObject {
	return t.objects
}

func newTableRender(game *Game) *tableRender {
	render := &tableRender{}
	render.game = game

	render.bg = canvas.NewRectangle(color.RGBA{0x07, 0x63, 0x24, 0xff})
	render.deck = newCard(faces.Back)

	render.pile1 = newCard(nil)
	render.pile2 = newCard(nil)
	render.pile3 = newCard(nil)

	render.space1 = newCard(faces.Space)
	render.space2 = newCard(faces.Space)
	render.space3 = newCard(faces.Space)
	render.space4 = newCard(faces.Space)

	render.stack1 = newStackRender()
	render.stack2 = newStackRender()
	render.stack3 = newStackRender()
	render.stack4 = newStackRender()
	render.stack5 = newStackRender()
	render.stack6 = newStackRender()
	render.stack7 = newStackRender()

	render.objects = []fyne.CanvasObject{render.bg, render.deck, render.pile1, render.pile2, render.pile3,
		render.space1, render.space2, render.space3, render.space4}

	render.objects = append(render.objects, render.stack1.cards[0:]...)
	render.objects = append(render.objects, render.stack2.cards[0:]...)
	render.objects = append(render.objects, render.stack3.cards[0:]...)
	render.objects = append(render.objects, render.stack4.cards[0:]...)
	render.objects = append(render.objects, render.stack5.cards[0:]...)
	render.objects = append(render.objects, render.stack6.cards[0:]...)
	render.objects = append(render.objects, render.stack7.cards[0:]...)

	return render
}

type stackRender struct {
	cards [13]fyne.CanvasObject
}

func (s *stackRender) Layout(pos fyne.Position, size fyne.Size) {
	top := pos.Y
	for i, _ := range s.cards {
		s.cards[i].Move(fyne.NewPos(pos.X, top))

		top += overlap
	}
}

func (s *stackRender) Refresh(stack Stack) {
	for i, _ := range s.cards {
		if i < len(stack.Cards)-1 {
			s.cards[i].(*canvas.Image).File = faces.Back.CachePath()
		} else if i == len(stack.Cards)-1 {
			s.cards[i].(*canvas.Image).File = stack.Cards[i].Face().CachePath()
		}
	}
}

func newStackRender() *stackRender {
	r := &stackRender{}
	for i, _ := range r.cards {
		r.cards[i] = newCard(nil)
	}

	return r
}
