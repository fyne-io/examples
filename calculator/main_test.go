package calculator

import (
	"testing"

	"fyne.io/fyne/test"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	calc := newCalculator()
	calc.loadUI(test.NewApp())

	test.Tap(calc.buttons["1"])
	test.Tap(calc.buttons["+"])
	test.Tap(calc.buttons["1"])
	test.Tap(calc.buttons["="])

	assert.Equal(t, "2", calc.output.Text)
}

func TestClear(t *testing.T) {
	calc := newCalculator()
	calc.loadUI(test.NewApp())

	test.Tap(calc.buttons["1"])
	test.Tap(calc.buttons["2"])
	test.Tap(calc.buttons["C"])

	assert.Equal(t, "", calc.output.Text)
}

func TestKeyboard(t *testing.T) {
	calc := newCalculator()
	calc.loadUI(test.NewApp())

	test.TypeOnCanvas(calc.window.Canvas(), "1+1")
	assert.Equal(t, "1+1", calc.output.Text)

	test.TypeOnCanvas(calc.window.Canvas(), "=")
	assert.Equal(t, "2", calc.output.Text)
}
