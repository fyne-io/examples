package fsviewer

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"fyne.io/fyne/storage"
	"fyne.io/fyne/test"

	"github.com/stretchr/testify/assert"
)

func Test_newFileTree(t *testing.T) {
	test.NewApp()

	tempDir, err := ioutil.TempDir("", "test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)
	err = os.MkdirAll(path.Join(tempDir, "A"), os.ModePerm)
	assert.NoError(t, err)
	err = os.MkdirAll(path.Join(tempDir, "B"), os.ModePerm)
	assert.NoError(t, err)
	err = ioutil.WriteFile(path.Join(tempDir, "B", "C"), []byte("c"), os.ModePerm)
	assert.NoError(t, err)

	root := storage.NewURI("file://" + tempDir)
	tree := newFileTree(root, nil, nil)
	tree.OpenAllBranches()

	assert.True(t, tree.IsBranchOpen(root.String()))
	b1, err := storage.Child(root, "A")
	assert.NoError(t, err)
	assert.True(t, tree.IsBranchOpen(b1.String()))
	b2, err := storage.Child(root, "B")
	assert.NoError(t, err)
	assert.True(t, tree.IsBranchOpen(b2.String()))
	l1, err := storage.Child(b2, "C")
	assert.NoError(t, err)
	assert.False(t, tree.IsBranchOpen(l1.String()))
}
