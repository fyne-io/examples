package sudoku

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/fyne-io/examples/img/icon"

	"github.com/andydotxyz/sudokgo"
)

type gui struct {
	sudoku *sudokgo.Sudoku
	cells  [sudokgo.RowSize][sudokgo.RowSize]*widget.Entry
}

// Show starts a new sudoku UI
func Show(a fyne.App) {
	game := sudokgo.NewSudoku()
	gui := newGUI(game)

	content := gui.LoadUI()
	go gui.generate()

	win := a.NewWindow("Sudoku")
	win.SetIcon(icon.SudokuBitmap)
	win.SetContent(content)
	win.Show()
}

func (g *gui) newCell(x, y int) *widget.Entry {
	entry := &widget.Entry{}
	g.cells[x][y] = entry
	return entry
}

func (g *gui) newRow(c *fyne.Container, y int) {
	for i := 0; i < sudokgo.RowSize; i++ {
		c.AddObject(g.newCell(i, y))
	}
}

func (g *gui) refresh() {
	for x := 0; x < sudokgo.RowSize; x++ {
		for y := 0; y < sudokgo.RowSize; y++ {
			entry := g.cells[x][y]
			value := g.sudoku.Grid[x][y]
			if value == -1 {
				entry.SetText("")
			} else {
				entry.SetText(fmt.Sprintf("%d", value))
			}
		}
	}
}

func (g *gui) loadButtons() fyne.CanvasObject {
	random := widget.NewButton("Random", func() {
		g.generate()
	})
	solve := widget.NewButton("Solve", func() {
		g.solve()
	})
	reset := widget.NewButton("Reset", func() {
		g.reset()
	})

	return widget.NewHBox(random, layout.NewSpacer(), solve, reset)
}

func (g *gui) LoadUI() fyne.CanvasObject {
	cells := fyne.NewContainerWithLayout(layout.NewGridLayout(sudokgo.RowSize))
	for i := 0; i < sudokgo.RowSize; i++ {
		g.newRow(cells, i)
	}

	buttons := g.loadButtons()

	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, buttons, nil, nil),
		buttons, cells)
}

func (g *gui) generate() {
	g.sudoku.Generate(sudokgo.ScoreEasy)
	g.refresh()
}

func (g *gui) solve() {
	g.sudoku.Solve()
	g.refresh()
}

func (g *gui) reset() {
	g.refresh()
}

func newGUI(s *sudokgo.Sudoku) *gui {
	ret := &gui{sudoku: s}

	return ret
}
