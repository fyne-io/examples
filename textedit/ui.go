package textedit

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type textEdit struct {
	cursorRow, cursorCol *widget.Label
	entry                *widget.Entry
}

func (e *textEdit) updateStatus() {
	e.cursorRow.SetText(fmt.Sprintf("%d", e.entry.CursorRow+1))
	e.cursorCol.SetText(fmt.Sprintf("%d", e.entry.CursorColumn+1))
}

// Show loads a new text editor
func Show(app fyne.App) {
	window := app.NewWindow("Text Editor")

	entry := widget.NewMultiLineEntry()
	toolbar := widget.NewToolbar(widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
		entry.SetText("")
	}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ContentCutIcon(), func() {
			entry.TypedShortcut(&fyne.ShortcutCut{Clipboard: window.Clipboard()})
		}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {
			entry.TypedShortcut(&fyne.ShortcutCopy{Clipboard: window.Clipboard()})
		}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {
			entry.TypedShortcut(&fyne.ShortcutPaste{Clipboard: window.Clipboard()})
		}))

	cursorRow := widget.NewLabel("1")
	cursorCol := widget.NewLabel("1")

	status := widget.NewHBox(layout.NewSpacer(),
		widget.NewLabel("Cursor Row:"), cursorRow,
		widget.NewLabel("Col:"), cursorCol)
	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, status, nil, nil),
		toolbar, status, widget.NewScrollContainer(entry))

	window.SetContent(content)
	window.Resize(fyne.NewSize(480, 320))

	editor := &textEdit{
		cursorRow: cursorRow,
		cursorCol: cursorCol,
		entry:     entry,
	}
	editor.entry.OnCursorChanged = func() {
		editor.updateStatus()
	}

	window.Show()
}
