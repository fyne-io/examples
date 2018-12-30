package calculator

import (
	"testing"

	"github.com/fyne-io/fyne/test"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	calc := newCalculator()
	calc.loadUI(test.NewApp())

	test.Click(calc.buttons["1"])
	test.Click(calc.buttons["+"])
	test.Click(calc.buttons["1"])
	test.Click(calc.buttons["="])

	assert.Equal(t, "2", calc.output.Text)
}

func TestClear(t *testing.T) {
	calc := newCalculator()
	calc.loadUI(test.NewApp())

	test.Click(calc.buttons["1"])
	test.Click(calc.buttons["2"])
	test.Click(calc.buttons["C"])

	assert.Equal(t, "", calc.output.Text)
}
