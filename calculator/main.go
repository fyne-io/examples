package calculator

import (
	"fmt"
	"log"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/Knetic/govaluate"
)

type calc struct {
	equation  string
	functions map[string]func()

	output  *widget.Label
	buttons map[string]*widget.Button
	window  fyne.Window
}

func (c *calc) display(newtext string) {
	c.equation = newtext
	c.output.SetText(newtext)
}

func (c *calc) character(char rune) {
	c.display(c.equation + string(char))
}

func (c *calc) digit(d int) {
	r := rune(d)
	r += '0'
	c.character(r)
}

func (c *calc) clear() {
	c.display("")
}

func (c *calc) evaluate() {
	expression, err := govaluate.NewEvaluableExpression(c.output.Text)
	if err == nil {
		result, err := expression.Evaluate(nil)
		if err == nil {
			c.display(strconv.FormatFloat(result.(float64), 'f', -1, 64))
		}
	}

	if err != nil {
		log.Println("Error in calculation", err)
		c.display("error")
	}

	c.equation = ""
}

func (c *calc) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button

	return button
}

func (c *calc) digitButton(number int) *widget.Button {
	str := fmt.Sprintf("%d", number)
	action := func() {
		c.digit(number)
	}
	c.functions[str] = action

	return c.addButton(str, action)
}

func (c *calc) charButton(char rune) *widget.Button {
	action := func() {
		c.character(char)
	}
	c.functions[string(char)] = action

	return c.addButton(string(char), action)
}

func (c *calc) typedRune(r rune) {
	if r == '=' {
		c.evaluate()
		return
	} else if r == 'c' {
		c.clear()
		return
	}

	action := c.functions[string(r)]
	if action != nil {
		action()
	}
}

func (c *calc) typedKey(ev *fyne.KeyEvent) {
	if ev.Name == fyne.KeyReturn || ev.Name == fyne.KeyEnter {
		c.evaluate()
		return
	}
}

func (c *calc) loadUI(app fyne.App) {
	c.output = widget.NewLabel("")
	c.output.Alignment = fyne.TextAlignTrailing
	c.output.TextStyle.Monospace = true
	equals := c.addButton("=", func() {
		c.evaluate()
	})
	equals.Style = widget.PrimaryButton

	c.window = app.NewWindow("Calc")
	c.window.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		c.output,
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.charButton('+'),
			c.charButton('-'),
			c.charButton('*'),
			c.charButton('/')),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.digitButton(7),
			c.digitButton(8),
			c.digitButton(9),
			c.addButton("C", func() {
				c.clear()
			})),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.digitButton(4),
			c.digitButton(5),
			c.digitButton(6),
			c.charButton('(')),
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			c.digitButton(1),
			c.digitButton(2),
			c.digitButton(3),
			c.charButton(')')),
		fyne.NewContainerWithLayout(layout.NewGridLayout(2),
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				c.digitButton(0),
				c.charButton('.')),
			equals)),
	)

	c.window.Canvas().SetOnTypedRune(c.typedRune)
	c.window.Canvas().SetOnTypedKey(c.typedKey)
	c.window.Show()
}

func newCalculator() *calc {
	c := &calc{}
	c.functions = make(map[string]func())
	c.buttons = make(map[string]*widget.Button)

	return c
}

// Show loads a calculator example window for the specified app context
func Show(app fyne.App) {
	c := newCalculator()
	c.loadUI(app)
}
