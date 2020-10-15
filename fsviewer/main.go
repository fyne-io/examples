package fsviewer

import (
	"fmt"
	"net/url"
	"os"
	"sort"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Show loads a file system viewer example window for the specified app context
func Show(app fyne.App) {
	dir, err := os.UserHomeDir()
	if err != nil {
		fyne.LogError("Could not get user home directory", err)
		dir, err = os.Getwd()
		if err != nil {
			fyne.LogError("Could not get current working directory", err)
		}
	}
	if dir == "" {
		// Can't get any useful directory
		return
	}
	tree := newFileTree(
		storage.NewURI("file://"+dir),
		func(uid fyne.URI) bool {
			return true // No filtering
		},
		func(a, b fyne.URI) bool {
			return true // No sorting
		},
	)
	window := app.NewWindow("File System Viewer")
	// TODO window.SetIcon(icon.FileSystemBitmap)
	window.SetContent(tree)
	window.Show()
}

func newFileTree(root fyne.URI, filter func(fyne.URI) bool, sorter func(fyne.URI, fyne.URI) bool) *widget.Tree {
	tree := &widget.Tree{
		Root: root.String(),
		IsBranch: func(uid string) bool {
			_, err := storage.ListerForURI(storage.NewURI(uid))
			return err == nil
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			var icon fyne.CanvasObject
			if branch {
				icon = widget.NewIcon(nil)
			} else {
				icon = widget.NewFileIcon(nil)
			}
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), icon, widget.NewLabel("Template Object"))
		},
	}
	tree.ChildUIDs = func(uid string) (c []string) {
		luri, err := storage.ListerForURI(storage.NewURI(uid))
		if err != nil {
			fyne.LogError("Unable to get lister for "+uid, err)
		} else {
			uris, err := luri.List()
			if err != nil {
				fyne.LogError("Unable to list "+luri.String(), err)
			} else {
				// Filter URIs
				var us []fyne.URI
				for _, u := range uris {
					if filter == nil || filter(u) {
						us = append(us, u)
					}
				}
				// Sort URIs
				if sorter != nil {
					sort.Slice(us, func(i, j int) bool {
						return sorter(us[i], us[j])
					})
				}
				// Convert to Strings
				for _, u := range us {
					c = append(c, u.String())
				}
			}
		}
		return
	}
	tree.UpdateNode = func(uid string, branch bool, node fyne.CanvasObject) {
		uri := storage.NewURI(uid)
		c := node.(*fyne.Container)
		if branch {
			var r fyne.Resource
			if tree.IsBranchOpen(uid) {
				// Set open folder icon
				r = theme.FolderOpenIcon()
			} else {
				// Set folder icon
				r = theme.FolderIcon()
			}
			c.Objects[0].(*widget.Icon).SetResource(r)
		} else {
			// Set file uri to update icon
			c.Objects[0].(*widget.FileIcon).SetURI(uri)
		}
		l := c.Objects[1].(*widget.Label)
		if tree.Root == uid {
			l.SetText(uid)
		} else {
			l.SetText(uri.Name())
		}
	}
	tree.OnSelectionChanged = func(uid string) {
		fmt.Println("TreeNodeSelected:", uid)
		u, err := url.Parse(uid)
		if err != nil {
			fyne.LogError("Failed to parse url", err)
		} else {
			err := fyne.CurrentApp().OpenURL(u)
			if err != nil {
				fyne.LogError("Failed to open url", err)
			}
		}
	}
	tree.ExtendBaseWidget(tree)
	return tree
}
