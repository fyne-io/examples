package life

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoard_CountNeighbours(t *testing.T) {
	b := &board{}
	b.createGrid(3, 3)
	b.cells[0][0] = true
	b.cells[0][2] = true
	b.cells[2][0] = true
	b.cells[2][2] = true

	assert.Equal(t, 4, b.countNeighbours(1, 1))
}

func TestBoard_CountNeighbours_Corner(t *testing.T) {
	b := &board{}
	b.createGrid(2, 2)
	b.cells[0][1] = true
	b.cells[1][0] = true
	b.cells[1][1] = true

	assert.Equal(t, 3, b.countNeighbours(0, 0))
}

func TestBoard_CountNeighbours_IgnoresMiddle(t *testing.T) {
	b := &board{}
	b.createGrid(3, 3)
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			b.cells[y][x] = true
		}
	}

	assert.Equal(t, 8, b.countNeighbours(1, 1))
}

func TestBoard_CreateGrid(t *testing.T) {
	b := &board{}
	b.createGrid(5, 3)

	assert.Equal(t, 3, len(b.cells))
	assert.Equal(t, 5, len(b.cells[0]))
}

func TestBoard_EnsureGridSize(t *testing.T) {
	b := &board{}
	b.createGrid(2, 2)
	b.cells[1][1] = true

	b.ensureGridSize(4, 4)
	assert.Equal(t, 4, len(b.cells))
	assert.Equal(t, 4, len(b.cells[0]))
	assert.True(t, b.cells[1][1])
}

func TestBoard_Generatoin(t *testing.T) {
	b := &board{}
	b.createGrid(5, 5)
	assert.Equal(t, 0, b.generation)

	b.nextGen()
	assert.Equal(t, 1, b.generation)
}
