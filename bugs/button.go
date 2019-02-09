package bugs

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type bugRenderer struct {
	icon  *canvas.Image
	label *canvas.Text

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

func (b *bugRenderer) BackgroundColor() color.Color {
	return theme.ButtonColor()
}

func (b *bugRenderer) Refresh() {
	b.label.Text = b.button.text

	b.icon.Hidden = b.button.icon == nil
	if b.button.icon != nil {
		b.icon.Resource = b.button.icon
	}

	b.Layout(b.button.Size())
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

	size   fyne.Size
	pos    fyne.Position
	hidden bool
}

func (b *bugButton) Size() fyne.Size {
	return b.size
}

func (b *bugButton) Resize(size fyne.Size) {
	b.size = size

	if widget.Renderer(b) != nil {
		widget.Renderer(b).Layout(size)
	}
}

func (b *bugButton) Position() fyne.Position {
	return b.pos
}

func (b *bugButton) Move(pos fyne.Position) {
	b.pos = pos
}

func (b *bugButton) MinSize() fyne.Size {
	return widget.Renderer(b).MinSize()
}

func (b *bugButton) Visible() bool {
	return !b.hidden
}

func (b *bugButton) Show() {
	b.hidden = false
}

func (b *bugButton) Hide() {
	b.hidden = true
}

// Tapped is called when a regular tap is reported
func (b *bugButton) Tapped(ev *fyne.PointEvent) {
	b.tap(true)
}

// TappedSecondary is called when an alternative tap is reported
func (b *bugButton) TappedSecondary(ev *fyne.PointEvent) {
	b.tap(false)
}

func (b *bugButton) CreateRenderer() fyne.WidgetRenderer {
	text := canvas.NewText(b.text, theme.TextColor())
	text.Alignment = fyne.TextAlignCenter
	text.TextStyle.Bold = true

	icon := canvas.NewImageFromResource(b.icon)

	objects := []fyne.CanvasObject{
		text,
		icon,
	}

	return &bugRenderer{icon, text, objects, b}
}

// SetText allows the button label to be changed
func (b *bugButton) SetText(text string) {
	b.text = text

	widget.Refresh(b)
}

// SetIcon updates the icon on a label - pass nil to hide an icon
func (b *bugButton) SetIcon(icon fyne.Resource) {
	b.icon = icon

	widget.Refresh(b)
}

// newButton creates a new button widget with the specified label, themed icon and tap handler
func newButton(label string, icon fyne.Resource, tap func(bool)) *bugButton {
	return &bugButton{text: label, icon: icon, tap: tap}
}
