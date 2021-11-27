package widgets

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

//
// This is an example widget that uses an Animator to animate the display.
// Please use this code as a starting point for your own widget.
//

//
// These do not add to the Widget they are here to
// indicate conformance to specific interfaces.
// If you break the widget implementation of an interface
// the respective line below will indicate an error.
// This is really usefull for widget development
//   Note that for the Widget itself it will conform to
//   fyne.Widget because of widget.BaseWidget in the struct.
//
var _ fyne.WidgetRenderer = (*widgetAnimateRenderer)(nil)
var _ fyne.CanvasObject = (*WidgetAnimate)(nil)
var _ fyne.Widget = (*WidgetAnimate)(nil)
var _ fyne.Tappable = (*WidgetAnimate)(nil)
var _ fyne.SecondaryTappable = (*WidgetAnimate)(nil)
var _ fyne.Disableable = (*WidgetAnimate)(nil)

//
// widgetAnimator:
//
// Managing the animator can be a pain. This structure contains
// all of the state necessary for this particular animator.
//
// Once set up the widget can just call animate()
//
// The animator is created by the Renderer because the Renderer owns
// the display objects that require animating.
// This creates an issue because the widget is the component that
// must start the animator.
//
type widgetAnimator struct {
	animator *fyne.Animation // The fyne animator
	// The rest of the items are to control the animation over time.
	animDelta      float32           // The amount of movement accumulated for each animation
	animDeltaReset float32           // Uset to reset the Delta so we can start again
	animDone       bool              // Dont start it if already running
	animObj        fyne.CanvasObject // The object that needs animating
}

func newWidgetAnimator(animObj fyne.CanvasObject, delta float32) *widgetAnimator {
	an := &widgetAnimator{
		animator:       &fyne.Animation{Duration: canvas.DurationStandard},
		animObj:        animObj,
		animDone:       true,
		animDeltaReset: delta,
		animDelta:      delta,
	}
	// Set the function called by the Animator
	an.animator.Tick = an.animTick
	return an
}

//
// Called for each tick of the Animator.
//    f goes from 0.0 to 1.0. This is defined by the animator.
//    When f = 0.5 change the direction of the animation (delta)
// 	  When f = 1.0 we have finished.
//    f is rounded to 1 decimal place (ff) to make the comparison reliable
//
func (an *widgetAnimator) animTick(f float32) {
	// Round f to 2 dp
	ff := math.Round(float64(f)*10.0) / 10
	if ff == 0.5 {
		// Change direction
		an.animDelta = an.animDelta * -1
	}
	if ff == 1.0 {
		// Finished. Stop the animation and set animDone true
		an.animator.Stop()
		an.animDone = true
	}
	// Calk the new size of the circle
	ns := fyne.NewSize(an.animObj.Size().Width-an.animDelta, an.animObj.Size().Height-an.animDelta)
	//
	// Defensive check that width and height are above 0 otherwise a panic happens.
	// This may happen if too many calls to animTick take place or step size (animDelta) is too big
	//
	if ns.Width > 0 && ns.Height > 0 {
		// Resize, Move and re-draw (Refresh)
		an.animObj.Resize(ns)
		an.animObj.Move(fyne.NewPos(an.animObj.Position().X+an.animDelta/2, an.animObj.Position().Y+an.animDelta/2))
		an.animObj.Refresh()
	}
}

//
// Start the animation as long as it is not already running
//
func (an *widgetAnimator) animate() {
	if an.animDone {
		an.animDone = false              // Prevent re-run
		an.animDelta = an.animDeltaReset // Reset the delta
		an.animator.Start()              // Start animation
	}
}

//
// widgetAnimateRenderer:
//
// The Renderer owns all of the canvas (renderable) objects and knows it's size.
//
// This is why the Renderer must create the Animator
//
// Once created the Renderer is passed back to the fyne app and cannot be referenced
// by the widget.
//
type widgetAnimateRenderer struct {
	minSize fyne.Size
	size    fyne.Size
	circ1   *canvas.Circle // This is the object to be animated
	rect1   *canvas.Rectangle
}

//
// Creare the Renderer (and the Animator). Both are passed back to the widget.
//
func newWidgetAnimateRenderer(min fyne.Size) (*widgetAnimateRenderer, *widgetAnimator) {
	// Create the renderer with Canvas objects and set minSize and size
	r := &widgetAnimateRenderer{
		minSize: min, size: min,
		circ1: &canvas.Circle{
			StrokeWidth: 2,
			StrokeColor: color.White,
			FillColor:   color.Transparent,
		}, rect1: &canvas.Rectangle{
			StrokeWidth: 1,
			StrokeColor: color.White,
			FillColor:   color.Transparent,
		},
	}
	// Create the Animator with the circle and the animation step size.
	return r, newWidgetAnimator(r.circ1, 4)
}

//
// Layout: From the WidgetRenderer interface.
// Layout the circle and the rectangle for a given size
//
func (r *widgetAnimateRenderer) Layout(s fyne.Size) {
	// Defensive check that width and height are above 0 otherwise a panic happens.
	if s.Width <= 1 || s.Height <= 1 {
		return
	}
	r.size = s
	r.circ1.Resize(s)
	r.rect1.Resize(s)
}

//
// MinSize: From the WidgetRenderer interface.
// This is the minimum size of the widget
//
func (r *widgetAnimateRenderer) MinSize() fyne.Size {
	return r.minSize
}

//
// From the WidgetRenderer interface.
// Does not seem to be called! It is required for interface complience
//
func (r *widgetAnimateRenderer) Refresh() {
}

//
// From the WidgetRenderer interface.
// Return a list of CanvasObjects that will require display (rendering)
// The order is critical. The last object in the list is the last to be drawn
//
func (r *widgetAnimateRenderer) Objects() []fyne.CanvasObject {
	o := make([]fyne.CanvasObject, 0)
	o = append(o, r.rect1, r.circ1)
	return o
}

//
// From the WidgetRenderer interface.
// Called when the rendered is destroyed and the memory released.
//  This is where you clean up your mess!
//  Nothing to do at the moment!
//
func (r *widgetAnimateRenderer) Destroy() {
}

//
// The state of the widget.
//
type WidgetAnimate struct {
	widget.BaseWidget // WidgetAnimate inherits the interface from BaseWidget
	minSize           fyne.Size
	animator          *widgetAnimator        // The animator so it can be syarted when the widget is clicked
	disabled          bool                   // Disabled. Will not animate if true
	tapped            func(*fyne.PointEvent) // Called when the widget is clicked unless disabled is true
}

//
// Create the widget and extend the base widget so we can inherit it's behaviour.
//
func NewWidgetAnimate(W, H float32, tapped func(*fyne.PointEvent)) *WidgetAnimate {
	w := &WidgetAnimate{
		disabled: false,
		tapped:   tapped,
		minSize:  fyne.Size{Width: W, Height: H},
	}
	w.ExtendBaseWidget(w)
	return w
}

//
// TappedSecondary: From the fyne.SecondaryTappable interface
// Calls the tapped function (if defined and not disabled) after starting the animator
//
func (w *WidgetAnimate) TappedSecondary(pe *fyne.PointEvent) {
	if w.tapped != nil && !w.disabled {
		if w.animator != nil {
			// Start the animator if it is defined
			w.animator.animate()
		}
		// Call the tapped function passed in when the widget is created
		w.tapped(pe)
	}
}

//
// Tapped: From the fyne.Tappable interface
// Calls the tapped function (if defined and not disabled) after starting the animator
//
func (w *WidgetAnimate) Tapped(pe *fyne.PointEvent) {
	if w.tapped != nil && !w.disabled {
		if w.animator != nil {
			// Start the animator if it is defined
			w.animator.animate()
		}
		// Call the tapped function passed in when the widget is created
		w.tapped(pe)
	}
}

//
// The fyne.Disableable interface defines these three functions.
//
func (w *WidgetAnimate) Enable()        { w.disabled = false }
func (w *WidgetAnimate) Disable()       { w.disabled = true }
func (w *WidgetAnimate) Disabled() bool { return w.disabled }

//
// Called by fyne app when it needs the renderer.
//
// It also keeps the reference to the Animator created bt the Renderer.
//
func (w *WidgetAnimate) CreateRenderer() fyne.WidgetRenderer {
	r, a := newWidgetAnimateRenderer(w.MinSize())
	w.animator = a
	return r
}

//
// The minimum size of the widget.
// Set by the W and H parameters when created
//
func (w *WidgetAnimate) MinSize() fyne.Size {
	return w.minSize
}

//
// Pass Resize events to the Base Widget
//
func (w *WidgetAnimate) Resize(s fyne.Size) {
	w.BaseWidget.Resize(s)
}
