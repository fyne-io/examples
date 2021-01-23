package textedit

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type textEdit struct {
	cursorRow, cursorCol *widget.Label
	entry                *widget.Entry
	window               fyne.Window
}

func (e *textEdit) updateStatus() {
	e.cursorRow.SetText(fmt.Sprintf("%d", e.entry.CursorRow+1))
	e.cursorCol.SetText(fmt.Sprintf("%d", e.entry.CursorColumn+1))
}

func (e *textEdit) cut() {
	e.entry.TypedShortcut(&fyne.ShortcutCut{Clipboard: e.window.Clipboard()})
}

func (e *textEdit) copy() {
	e.entry.TypedShortcut(&fyne.ShortcutCopy{Clipboard: e.window.Clipboard()})
}

func (e *textEdit) paste() {
	e.entry.TypedShortcut(&fyne.ShortcutPaste{Clipboard: e.window.Clipboard()})
}

func (e *textEdit) buildToolbar() *widget.Toolbar {
	return widget.NewToolbar(widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
		e.entry.SetText("")
	}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			e.cut()
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			e.copy()
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			e.paste()
		}))
}

// Show loads a new text editor
func Show(win fyne.Window) fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()
	cursorRow := widget.NewLabel("1")
	cursorCol := widget.NewLabel("1")

	editor := &textEdit{
		cursorRow: cursorRow,
		cursorCol: cursorCol,
		entry:     entry,
		window:    win,
	}

	toolbar := editor.buildToolbar()
	status := container.NewHBox(layout.NewSpacer(),
		widget.NewLabel("Cursor Row:"), cursorRow,
		widget.NewLabel("Col:"), cursorCol)
	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, status, nil, nil),
		toolbar, status, container.NewScroll(entry))

	editor.entry.OnCursorChanged = func() {
		editor.updateStatus()
	}

	return content
}
