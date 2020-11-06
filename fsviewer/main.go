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
		func(id fyne.URI) bool {
			return true // No filtering
		},
		func(a, b fyne.URI) bool {
			return true // No sorting
		},
	)
	window := app.NewWindow("File System Viewer")
	// TODO window.SetIcon(icon.FileSystemBitmap)
	window.SetContent(tree)
	window.Resize(fyne.NewSize(600, 400))
	window.Show()
}

func newFileTree(root fyne.URI, filter func(fyne.URI) bool, sorter func(fyne.URI, fyne.URI) bool) *widget.Tree {
	uriCache := make(map[widget.TreeNodeID]fyne.URI)
	luriCache := make(map[widget.TreeNodeID]fyne.ListableURI)
	tree := &widget.Tree{
		Root: root.String(),
		IsBranch: func(id widget.TreeNodeID) bool {
			if _, ok := luriCache[id]; ok {
				return true
			}
			uri, ok := uriCache[id]
			if !ok {
				uri = storage.NewURI(id)
				uriCache[id] = uri
			}
			if luri, err := storage.ListerForURI(uri); err == nil {
				luriCache[id] = luri
				return true
			}
			return false
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
	tree.ChildUIDs = func(id widget.TreeNodeID) (c []string) {
		luri, ok := luriCache[id]
		if !ok {
			uri, ok := uriCache[id]
			if !ok {
				uri = storage.NewURI(id)
				uriCache[id] = uri
			}
			l, err := storage.ListerForURI(uri)
			if err != nil {
				fyne.LogError("Unable to get lister for "+id, err)
				return
			} else {
				luri = l
				luriCache[id] = l
			}
		}
		uris, err := luri.List()
		if err != nil {
			fyne.LogError("Unable to list "+luri.String(), err)
			return
		}
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
		return
	}
	tree.UpdateNode = func(id widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		uri, ok := uriCache[id]
		if !ok {
			uri = storage.NewURI(id)
			uriCache[id] = uri
		}
		c := node.(*fyne.Container)
		if branch {
			var r fyne.Resource
			if tree.IsBranchOpen(id) {
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
		if tree.Root == id {
			l.SetText(id)
		} else {
			l.SetText(uri.Name())
		}
	}
	tree.OnSelected = func(id widget.TreeNodeID) {
		fmt.Println("TreeNodeSelected:", id)
		u, err := url.Parse(id)
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
