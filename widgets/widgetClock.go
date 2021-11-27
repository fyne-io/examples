package widgets

import (
	"fmt"
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//
// This is an example widget that uses a time.Ticker to update the display.
// Please use this code as a starting point for your own widget.
//
const (
	PI_180_F64           = float64(math.Pi / 180.0) // Used to convert degrees to radians
	DEG_PER_TICK_INT     = 6                        // 60 clock hand positions on a 360 degree face. 6 degrees per pos
	HOUR_TO_60_INT       = 5                        // 5 clock positions per hour for the hour hand
	MIN_IN_5DEG_F64      = float64(5.0 / 60.0)      // positions between numerals for minutes of the hour hand
	SEC_HAND_LEN_F32     = 0.77                     // adjusted length of the second hand
	MIN_HAND_LEN_F32     = 0.65                     // adjusted length of the minute hand
	HOUR_HAND_LEN_F32    = 0.55                     // adjusted length of the hour hand
	CENTER_SIZE_FACT_F32 = 15.0                     // Radius of center circle scaled from radius of clock face
)

//
// These do not add to the Widget they are here to
// indicate conformance to specific interfaces.
// If you break the widget implementation of an interface
// the respective line below will indicate an error.
// This is really usefull for widget development
//   Note that for the Widget itself it will conform to
//   fyne.Widget because of widget.BaseWidget in the struct.
//
var _ fyne.Widget = (*WidgetClock)(nil)                 // Test WidgetClock is a fyne.Widget
var _ fyne.WidgetRenderer = (*widgetClockRenderer)(nil) // Test widgetClockRenderer is a fyne.fyne.WidgetRenderer
var _ fyne.CanvasObject = (*WidgetClock)(nil)           // Test WidgetClock is a fyne.CanvasObject
var _ fyne.Tappable = (*WidgetClock)(nil)               // Test WidgetClock is a fyne.Tappable
var _ fyne.SecondaryTappable = (*WidgetClock)(nil)      // Test WidgetClock is a fyne.SecondaryTappable

//
// The Renderer (fyne.WidgetRenderer interface) is responsible for ALL drawing, scaling, moving, animating etc.
// Animation here is where we move the clock hands. This is done with a time.NewTicker not a fyne.Animator.
//
// Note h,m,s would usually be the responsibility of the Widget. Each is state information.
//      However, as these are updated by the Ticker and the Ticker is in the Renderer (for simplicity)
//      it is more logical to store them here.
//
type widgetClockRenderer struct {
	widget     *WidgetClock      // Reference to the widget so the renderer can access clock state information
	size       fyne.Size         // The current size of the widget. Set when Layout is called
	background *canvas.Rectangle // A rectangle that forms the background
	centCirc   *canvas.Circle    // A circle at the center of the clock dispay
	sHand      *canvas.Line      // The second hand
	mHand      *canvas.Line      // The minute hand
	hHand      *canvas.Line      // The hour hand
	h, m, s    int               // The hour, minute and second currently shown.
}

func newWidgetClockRenderer(w *WidgetClock) *widgetClockRenderer {
	r := &widgetClockRenderer{
		widget: w,
		size:   w.MinSize(),
		background: &canvas.Rectangle{
			FillColor: w.BackgroundColor,
		},
	}
	// Init the image to fill and keep it's aspect ratio
	r.widget.Image.FillMode = canvas.ImageFillContain
	// Create the canvas objects according to the state of the widget
	r.mHand = canvas.NewLine(w.MHandColor)
	r.mHand.StrokeWidth = 3
	r.hHand = canvas.NewLine(w.HHandColor)
	r.hHand.StrokeWidth = 3
	r.sHand = canvas.NewLine(w.SHandColor)
	r.sHand.StrokeWidth = 1
	r.centCirc = canvas.NewCircle(w.CenterColor)
	// Update the time at the start
	r.updateTime()
	//
	// Start the bacground animation using the Ticker.
	//  One process per second.
	//
	tick := time.NewTicker(time.Second)
	go func() {
		// Run forever in the background.
		// This code is NOT run when the widget is constructed.
		// It is run after and then after each second
		for {
			// If running update the time
			if r.widget.running {
				r.updateTime()
			}
			// Lay everything out using the minimum size
			r.Layout(r.size)
			// Redraw the widget
			canvas.Refresh(r.widget)
			// Wait for the next tick
			<-tick.C
		}
	}()
	// Return the renderer
	return r
}

//
// Layout: From the WidgetRenderer interface.
// For the given size 's'. Layout (re-arrange) all of the visible components.
// This also updates the position of the clock hands.
//
func (r *widgetClockRenderer) Layout(s fyne.Size) {
	// Somtimes Layout is called with invalid or 0 width or height.
	// If invalid then do nothing
	if s.Width <= 1 || s.Height <= 1 {
		return
	}

	r.size = s // Save the size
	// Calc image size and position.
	shrink := r.widget.ImageShrink
	offset := shrink / 2
	imageSize := fyne.Size{Width: s.Width - shrink, Height: s.Height - shrink}
	// Set image size and position.
	r.widget.Image.Resize(imageSize)
	r.widget.Image.Move(fyne.Position{X: offset, Y: offset})

	// Set background size and position.
	r.background.Resize(s)
	r.background.FillColor = r.widget.BackgroundColor

	// Calc center circle size and position.
	imageCent := fyne.Position{X: (imageSize.Width / 2) + offset, Y: (imageSize.Height / 2) + offset}
	centCircW := imageSize.Width / CENTER_SIZE_FACT_F32
	centCircH := imageSize.Height / CENTER_SIZE_FACT_F32
	// Set center circle size and position.
	r.centCirc.Resize(fyne.Size{Width: centCircW, Height: centCircH})
	r.centCirc.Move(fyne.Position{X: imageCent.X - (centCircW / 2), Y: imageCent.Y - (centCircH / 2)})

	// Set colours of all components
	r.centCirc.FillColor = r.widget.CenterColor
	r.sHand.StrokeColor = r.widget.SHandColor
	r.mHand.StrokeColor = r.widget.MHandColor
	r.hHand.StrokeColor = r.widget.HHandColor

	// Calc the full length of the line used to the clock hands
	// Uses the lowest width or height of tthe image (radius of the clock face)
	var lineLen float32
	if imageSize.Width > imageSize.Height {
		lineLen = imageSize.Height / 2
	} else {
		lineLen = imageSize.Width / 2
	}
	// Update each line (clock hand)
	// HAND_LEN_F32 is used to scale each hand to different lengths
	r.updateClockHand(imageCent, lineLen*SEC_HAND_LEN_F32, r.s, r.sHand)
	r.updateClockHand(imageCent, lineLen*MIN_HAND_LEN_F32, r.m, r.mHand)
	r.updateClockHand(imageCent, lineLen*HOUR_HAND_LEN_F32, hourToClock60(r.h, r.m), r.hHand)
}

//
// From the WidgetRenderer interface.
// Return the minimum size of the widget. Get this from the widget!
//
func (r *widgetClockRenderer) MinSize() fyne.Size {
	return r.widget.MinSize()
}

//
// From the WidgetRenderer interface.
// Does not seem to be called! It is required for interface complience
//
func (r *widgetClockRenderer) Refresh() {}

//
// From the WidgetRenderer interface.
// Return a list of CanvasObjects that will require display (rendering)
// The order is critical. The last object in the list is the last to be drawn
//
func (r *widgetClockRenderer) Objects() []fyne.CanvasObject {
	o := make([]fyne.CanvasObject, 0)
	o = append(o, r.background, r.widget.Image, r.sHand, r.hHand, r.mHand, r.centCirc)
	return o
}

//
// From the WidgetRenderer interface.
// Called when the rendered is destroyed and the memory released.
//  This is where you clean up your mess!
//  Nothing to do at the moment!
//
func (r *widgetClockRenderer) Destroy() {
}

//
// Update the time using the system clock
//  Do not call this if the clock is not running
//
func (r *widgetClockRenderer) updateTime() {
	r.h = time.Now().Hour()
	r.m = time.Now().Minute()
	r.s = time.Now().Second()
}

//
// Detirmine the position for the hour hand.
//    The are 60 positions corresponding to 360 degrees on the clock face.
//    The hour hand takes 12 hours plus additional degrees for additional minutes.
//    This means the hour hand moves between the clock numerals
//
func hourToClock60(hr, min int) int {
	return ((hr % 12) * HOUR_TO_60_INT) + int(float64(min)*MIN_IN_5DEG_F64)
}

//
// The are 60 positions corresponding to 360 degrees on the clock face.
// Given the position and length we use trigonometry to derive the end point of the line
//    x = len * cos(theta)
//    y = len * sin(theta)
// The start point of the line is the center of the clock face .
//   Note the math API uses float64 but every thing else uses float32 so
//   some conversion must take place.
//   Radians (float64) are used by math Cos and Sine functions.
//
func (r *widgetClockRenderer) updateClockHand(cent fyne.Position, lineLen float32, position int, line *canvas.Line) {
	radians := (float64((position*DEG_PER_TICK_INT)-90) * PI_180_F64)
	xEnd := cent.X + lineLen*float32(math.Cos(radians))
	yExd := cent.Y + lineLen*float32(math.Sin(radians))
	line.Position1 = cent
	line.Position2 = fyne.Position{X: xEnd, Y: yExd}
}

//
// The state of the widget.
// The properties here can be updated dynamically to change the
// apperance of the clock.
//
type WidgetClock struct {
	widget.BaseWidget               // Inherit from BaseWidget
	running           bool          // Clock time is updating (hands moving)
	minSize           fyne.Size     // The minimum size of the widget
	Image             *canvas.Image // An image displayed above the background (clock face)
	ImageShrink       float32       // Make the image smaller by this amount (provides a boarder)
	HHandColor        color.Color   // The hour hand colour
	MHandColor        color.Color   // The minute hand colour
	SHandColor        color.Color   // The second hand colour
	CenterColor       color.Color   // The blob at the clock center colour
	BackgroundColor   color.Color   // The background colour
}

//
// Create a Clock Widget and Extend the BaseWidget
// Minimum parameters provided.
//   imageShrink: The amount to shrink the image (provides a boarder)
//	 image: The clock face image.
//   W, H: Are used to set the minimum Width and Height of the widget
//   running: Will start the clock if true.
//       Use SetRunning(b bool) to start or stop the clock after it is created
//
func NewWidgetClock(W, H, imageShrink float32, image *canvas.Image, running bool) *WidgetClock {
	w := &WidgetClock{
		ImageShrink: imageShrink,
		Image:       image,
		minSize:     fyne.Size{Width: W, Height: H},
		running:     running,
		// Define the default values other properties
		MHandColor:      theme.PrimaryColorNamed("green"),
		SHandColor:      theme.PrimaryColorNamed("red"),
		HHandColor:      theme.PrimaryColorNamed("blue"),
		CenterColor:     theme.PrimaryColorNamed("green"),
		BackgroundColor: color.White,
	}
	w.ExtendBaseWidget(w)
	return w
}

//
// Part of the widget interface. Provided a reference to the Renderer
// Do not hang on to the reference to the Renderer. fyne may stop using
// it and call again for a new one. fyne also caches the reference.
// The Renderer gets the widget state passed in as a reference.
//
func (w *WidgetClock) CreateRenderer() fyne.WidgetRenderer {
	return newWidgetClockRenderer(w)
}

//
// Start (true) or Stop (false) the clock hands.
// The clock hands are always updated (in case the widget is resized)
// This flag just freezes their positions by stopping the Renderer calling updateTime()
//
func (w *WidgetClock) SetRunning(b bool) {
	w.running = b
}

//
// Returns true if the clock is running.
//
func (w *WidgetClock) GetRunning() bool {
	return w.running
}

//
// The minimum size of the widget.
// Set by the W and H parameters when created
//
func (w *WidgetClock) MinSize() fyne.Size {
	return w.minSize
}

//
// Pass Resize events to the Base Widget
//
func (w *WidgetClock) Resize(s fyne.Size) {
	w.BaseWidget.Resize(s)
}

//
// From the fyne.SecondaryTappable interface
//   Will be called if the clock face is (right) clicked
//
func (w *WidgetClock) TappedSecondary(pe *fyne.PointEvent) {
	fmt.Println("TappedSecondary")
}

//
// From the fyne.Tappable interface
//   Will be called if the clock face is (left) clicked
//
func (w *WidgetClock) Tapped(pe *fyne.PointEvent) {
	fmt.Println("Tapped")
}
