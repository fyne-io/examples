package xkch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/fyne-io/fyne"
	"github.com/fyne-io/fyne/widget"
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
	entry      map[string]*widget.Entry
}

func (x *XKCD) newEntry(name string) *widget.Entry {
	w := widget.NewEntry()
	x.entry[name] = w
	return w
}

// NewXKCD returns a new xkcd app
func NewXKCD() *XKCD {
	rand.Seed(time.Now().UnixNano())
	return &XKCD{
		entry: make(map[string]*widget.Entry),
	}
}

// Submit will lookup the xkcd cartoon and do something useful with it
func (x *XKCD) Submit() {
	fmt.Println("Submitted")

	// Get the ID
	id, _ := strconv.Atoi(x.entry["num"].Text)
	if id == 0 {
		id = rand.Intn(2075)
		fmt.Println("Getting random ID", id)
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

// DataToScreen copies the data model to the screen
func (x *XKCD) DataToScreen() {
	myType := reflect.TypeOf(x).Elem()
	myValue := reflect.ValueOf(x).Elem()
	for i := 0; i < myType.NumField(); i++ {
		tag := myType.Field(i).Tag.Get("json")
		switch tag {
		case "": // not a display field
		case "img": // special field for images
			// TODO - load the image into the img widged
		case "num":
			v := myValue.Field(i).Int()
			x.entry[tag].SetText(fmt.Sprintf("%d", v))
		default:
			v := myValue.Field(i).String()
			if newline := strings.IndexAny(v, "\n.-,"); newline > -1 {
				v = v[:newline] + "..."
			}
			x.entry[tag].SetText(v)
		}
	}
}

// NewForm generates a new XKCD form
func (x *XKCD) NewForm(w fyne.Window) fyne.Widget {
	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: x.Submit,
	}
	tt := reflect.TypeOf(x).Elem()
	for i := 0; i < tt.NumField(); i++ {
		fld := tt.Field(i)
		tag := fld.Tag.Get("json")
		switch tag {
		case "": // not a display field
		case "img": // special field for images
		// TODO - build an image display field here
		default:
			form.Append(fld.Name, x.newEntry(tag))
		}
	}
	return form
}

// Show starts a new xkcd widget
func Show(app fyne.App) {
	x := NewXKCD()
	w := app.NewWindow("XKCD Viewer")
	w.SetContent(x.NewForm(w))
	w.Show()
}

/*
// Qt code
package main
​
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
​
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)
​
var data_struct struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
​
func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
​
	window := widgets.NewQMainWindow(nil, 0)
	widget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(widget)
​
	layout := widgets.NewQFormLayout(widget)
	layout.SetFieldGrowthPolicy(widgets.QFormLayout__AllNonFixedFieldsGrow)
​
	widgetmap := make(map[string]*widgets.QWidget)
	for i := 0; i < reflect.TypeOf(data_struct).NumField(); i++ {
		name := reflect.TypeOf(data_struct).Field(i).Tag.Get("json")
​
		if name != "img" {
			widgetmap[name] = widgets.NewQLineEdit(nil).QWidget_PTR()
​
			layout.AddRow3(name, widgetmap[name])
		} else {
			widgetmap[name] = widgets.NewQLineEdit(nil).QWidget_PTR()
			widgetmap[name+"_label"] = widgets.NewQLabel(nil, 0).QWidget_PTR()
​
			layout.AddRow3(name, widgetmap[name])
			layout.AddRow3(name+"_label", widgetmap[name+"_label"])
		}
	}
​
	button := widgets.NewQPushButton2("random xkcd", nil)
	layout.AddWidget(button)
	button.ConnectClicked(func(bool) {
		rand.Seed(time.Now().UnixNano())
​
		resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%v/info.0.json", rand.Intn(614)))
		if err != nil {
			return
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
​
		json.Unmarshal(data, &data_struct)
​
		for i := 0; i < reflect.TypeOf(data_struct).NumField(); i++ {
			name := reflect.TypeOf(data_struct).Field(i).Tag.Get("json")
​
			if name != "img" {
				switch reflect.ValueOf(data_struct).Field(i).Kind() {
				case reflect.String:
					widgets.NewQLineEditFromPointer(widgetmap[name].Pointer()).SetText(reflect.ValueOf(data_struct).Field(i).String())
				case reflect.Int:
					widgets.NewQLineEditFromPointer(widgetmap[name].Pointer()).SetText(strconv.Itoa(int(reflect.ValueOf(data_struct).Field(i).Int())))
				}
			} else {
				url := reflect.ValueOf(data_struct).Field(i).String()
​
				widgets.NewQLineEditFromPointer(widgetmap[name].Pointer()).SetText(url)
​
				resp, err := http.Get(url)
				if err != nil {
					return
				}
				defer resp.Body.Close()
				data, _ := ioutil.ReadAll(resp.Body)
​
				widgets.NewQLabelFromPointer(widgetmap[name+"_label"].Pointer()).SetPixmap(gui.QPixmap_FromImage(gui.QImage_FromData(string(data), len(data), ""), 0).Scaled2(400, 400, core.Qt__KeepAspectRatio, core.Qt__SmoothTransformation))
			}
		}
	})
​
	window.Show()
	widgets.QApplication_Exec()
}
*/
