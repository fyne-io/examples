package bugs

import "github.com/fyne-io/fyne"
import "github.com/fyne-io/fyne/canvas"
import "github.com/fyne-io/fyne/theme"

type bugRenderer struct {
	background *canvas.Rectangle
	icon       *canvas.Image
	label      *canvas.Text

	objects []fyne.CanvasObject
	button  *bugButton
}

const bugSize = 18

// MinSize calculates the minimum size of a bug button. A fixed amount.
func (b *bugRenderer) MinSize() fyne.Size {
	return fyne.NewSize(bugSize+theme.Padding()*2, bugSize+theme.Padding()*2)
}

// Layout the components of the widget
func (b *bugRenderer) Layout(size fyne.Size) {
	b.background.Resize(size)

	inner := size.Subtract(fyne.NewSize(theme.Padding()*2, theme.Padding()*2))
	b.label.TextSize = inner.Height - 3

	b.label.Resize(inner)
	b.label.Move(fyne.NewPos(theme.Padding(), theme.Padding()))
	b.icon.Resize(inner)
	b.icon.Move(fyne.NewPos(theme.Padding(), theme.Padding()))
}

// ApplyTheme is called when the bugButton may need to update it's look
func (b *bugRenderer) ApplyTheme() {
	b.label.Color = theme.TextColor()

	b.Refresh()
}

func (b *bugRenderer) Refresh() {
	b.label.Text = b.button.text

	b.icon.Hidden = b.button.icon == nil
	if b.button.icon != nil {
		b.icon.File = b.button.icon.CachePath()
	}

	b.Layout(b.button.CurrentSize())
	canvas.Refresh(b.button)
}

func (b *bugRenderer) Objects() []fyne.CanvasObject {
	return b.objects
}

// bugButton widget is a scalable button that has a text label and icon and triggers an event func when clicked
type bugButton struct {
	text string
	icon fyne.Resource

	tap func(bool)

	size     fyne.Size
	pos      fyne.Position
	hidden   bool
	renderer fyne.WidgetRenderer
}

func (b *bugButton) CurrentSize() fyne.Size {
	return b.size
}

func (b *bugButton) Resize(size fyne.Size) {
	b.size = size

	if b.renderer != nil {
		b.renderer.Layout(size)
	}
}

func (b *bugButton) CurrentPosition() fyne.Position {
	return b.pos
}

func (b *bugButton) Move(pos fyne.Position) {
	b.pos = pos
}

func (b *bugButton) MinSize() fyne.Size {
	return b.Renderer().MinSize()
}

func (b *bugButton) IsVisible() bool {
	return !b.hidden
}

func (b *bugButton) Show() {
	b.hidden = false
}

func (b *bugButton) Hide() {
	b.hidden = true
}

// OnMouseDown is called when a mouse down event is captured and triggers any tap handler
func (b *bugButton) OnMouseDown(ev *fyne.MouseEvent) {
	b.tap(ev.Button == fyne.LeftMouseButton)
}

func (b *bugButton) createRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(b.text, theme.TextColor())
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle.Bold = true

	icon := canvas.NewImageFromResource(b.icon)
	bg := canvas.NewRectangle(theme.ButtonColor())

	objects := []fyne.CanvasObject{
		bg,
		text,
		icon,
	}

	return &bugRenderer{bg, icon, text, objects, b}
}

// Renderer is a private method to Fyne which links this widget to it's renderer
func (b *bugButton) Renderer() fyne.WidgetRenderer {
	if b.renderer == nil {
		b.renderer = b.createRenderer()
	}

	return b.renderer
}

// SetText allows the button label to be changed
func (b *bugButton) SetText(text string) {
	b.text = text

	b.Renderer().Refresh()
}

// SetIcon updates the icon on a label - pass nil to hide an icon
func (b *bugButton) SetIcon(icon fyne.Resource) {
	b.icon = icon

	b.Renderer().Refresh()
}

// newButton creates a new button widget with the specified label, themed icon and tap handler
func newButton(label string, icon fyne.Resource, tap func(bool)) *bugButton {
	return &bugButton{text: label, icon: icon, tap: tap}
}
