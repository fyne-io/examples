package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/fyne-io/examples/img/icon"
)

// XKCD is an app to get xkcd images and display them
type XKCD struct {
	ID         int    `json:"num"`
	Title      string `json:"title"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Link       string `json:"link"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	News       string `json:"news"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`

	image   *canvas.Image
	iDEntry *widget.Entry
	labels  map[string]*widget.Label
}

func (x *XKCD) newLabel(name string) *widget.Label {
	w := widget.NewLabel("")
	x.labels[name] = w
	return w
}

// NewXKCD returns a new xkcd app
func NewXKCD() *XKCD {
	rand.Seed(time.Now().UnixNano())
	return &XKCD{
		labels: make(map[string]*widget.Label),
	}
}

// Submit will lookup the xkcd cartoon and do something useful with it
func (x *XKCD) Submit() {
	// Get the ID
	id, _ := strconv.Atoi(x.iDEntry.Text)
	if id == 0 {
		id = rand.Intn(2075)
	}

	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", id))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		data, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(data, x)
		x.DataToScreen()
	} else {
		fmt.Println("Error getting ID", id, resp.Status, resp.StatusCode)
	}
}

func (x *XKCD) downloadImage(url string) {
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	file, err := ioutil.TempFile(os.TempDir(), "xkcd.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	x.image.File = file.Name()
	canvas.Refresh(x.image)
}

// DataToScreen copies the data model to the screen
func (x *XKCD) DataToScreen() {
	myType := reflect.TypeOf(x).Elem()
	myValue := reflect.ValueOf(x).Elem()
	for i := 0; i < myType.NumField(); i++ {
		tag := myType.Field(i).Tag.Get("json")
		switch tag {
		case "": // not a display field
		case "img": // special field for images
			url := myValue.Field(i).String()

			go x.downloadImage(url)
		case "num":
			v := myValue.Field(i).Int()
			x.iDEntry.SetText(fmt.Sprintf("%d", v))
		default:
			v := myValue.Field(i).String()
			if newline := strings.IndexAny(v, "\n.-,"); newline > -1 {
				v = v[:newline] + "..."
			}
			x.labels[tag].SetText(v)
		}
	}
}

// NewForm generates a new XKCD form
func (x *XKCD) NewForm(w fyne.Window) fyne.Widget {
	form := &widget.Form{}
	tt := reflect.TypeOf(x).Elem()
	for i := 0; i < tt.NumField(); i++ {
		fld := tt.Field(i)
		tag := fld.Tag.Get("json")
		switch tag {
		case "": // not a display field
		case "img": // special field for images
			// we created this in the setup
		case "num": // special field for ID
			entry := widget.NewEntry()
			x.iDEntry = entry
			form.Append(fld.Name, entry)
		default:
			form.Append(fld.Name, x.newLabel(tag))
		}
	}
	return form
}

// Show starts a new xkcd widget
func Show(app fyne.App) {
	x := NewXKCD()
	w := app.NewWindow("XKCD Viewer")
	w.SetIcon(icon.XKCDBitmap)

	form := x.NewForm(w)
	submit := widget.NewButton("Submit", func() {
		x.Submit()
	})
	submit.Style = widget.PrimaryButton
	buttons := widget.NewHBox(
		widget.NewButton("Cancel", func() {
			w.Close()
		}),
		layout.NewSpacer(),
		widget.NewButton("Random", func() {
			x.iDEntry.Text = ""
			x.Submit()
		}),
		submit)
	x.image = &canvas.Image{FillMode: canvas.ImageFillOriginal}
	imageContainer := widget.NewScrollContainer(x.image)
	controlsContainer := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(buttons, form, nil, nil),
		buttons, form)
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewBorderLayout(controlsContainer, nil, nil, nil),
		controlsContainer, imageContainer))
	w.Show()
}
