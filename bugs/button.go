package bugs

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
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
	b.icon.Resize(inner)
	b.icon.Move(fyne.NewPos(theme.Padding(), theme.Padding()))

	textSize := size.Height * .67
	textMin := fyne.CurrentApp().Driver().RenderedTextSize(b.label.Text, textSize, fyne.TextStyle{Bold: true})

	b.label.TextSize = textSize
	b.label.Resize(fyne.NewSize(size.Width, textMin.Height))
	b.label.Move(fyne.NewPos(0, (size.Height-textMin.Height)/2))
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

func (b *bugRenderer) Destroy() {
}

// bugButton widget is a scalable button that has a text label and icon and triggers an event func when clicked
type bugButton struct {
	widget.BaseWidget
	text string
	icon fyne.Resource

	tap func(bool)
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
	icon.FillMode = canvas.ImageFillContain

	objects := []fyne.CanvasObject{
		text,
		icon,
	}

	return &bugRenderer{icon, text, objects, b}
}

// SetText allows the button label to be changed
func (b *bugButton) SetText(text string) {
	b.text = text

	b.Refresh()
}

// SetIcon updates the icon on a label - pass nil to hide an icon
func (b *bugButton) SetIcon(icon fyne.Resource) {
	b.icon = icon

	b.Refresh()
}

// newButton creates a new button widget with the specified label, themed icon and tap handler
func newButton(label string, icon fyne.Resource, tap func(bool)) *bugButton {
	button := &bugButton{text: label, icon: icon, tap: tap}
	button.ExtendBaseWidget(button)
	return button
}
