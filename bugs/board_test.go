package bugs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countMines(b *board) int {
	count := 0

	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			item := b.bugs[y][x]

			if item.bug {
				count++
			}
		}
	}

	return count
}

func TestBoard(t *testing.T) {
	b := newBoard(10, 10)

	assert.Equal(t, 10, b.width)

	b.load(0)
	assert.Equal(t, 10, len(b.bugs))
	assert.Equal(t, 10, len(b.bugs[0]))

	assert.Equal(t, 0, b.bugs[2][2].near)
	assert.Equal(t, false, b.bugs[2][2].shown)
}

func TestBoard_load(t *testing.T) {
	b := newBoard(10, 10)
	b.load(4)

	assert.Equal(t, 4, countMines(b))
}

func TestBoard_setMine(t *testing.T) {
	b := newBoard(3, 3)
	b.load(0)
	b.setMine(1, 1)

	assert.True(t, b.bugs[1][1].bug)

	assert.Equal(t, 1, b.bugs[0][0].near)
	assert.Equal(t, 1, b.bugs[0][1].near)
	assert.Equal(t, 1, b.bugs[0][2].near)
	assert.Equal(t, 1, b.bugs[1][0].near)
	assert.Equal(t, 1, b.bugs[1][2].near)
	assert.Equal(t, 1, b.bugs[2][0].near)
	assert.Equal(t, 1, b.bugs[2][1].near)
	assert.Equal(t, 1, b.bugs[2][2].near)
}

func TestBoard_setMines(t *testing.T) {
	b := newBoard(3, 4)
	b.load(0)
	b.setMine(1, 1)
	b.setMine(2, 1)

	assert.Equal(t, 2, countMines(b))
	assert.True(t, b.bugs[1][1].bug)
	assert.True(t, b.bugs[1][2].bug)

	assert.Equal(t, 1, b.bugs[0][0].near)
	assert.Equal(t, 2, b.bugs[0][1].near)
	assert.Equal(t, 2, b.bugs[0][2].near)
	assert.Equal(t, 1, b.bugs[0][3].near)
	assert.Equal(t, 1, b.bugs[1][0].near)
	assert.Equal(t, 1, b.bugs[1][3].near)
	assert.Equal(t, 1, b.bugs[2][0].near)
	assert.Equal(t, 2, b.bugs[2][1].near)
	assert.Equal(t, 2, b.bugs[2][2].near)
	assert.Equal(t, 1, b.bugs[2][3].near)
}

func TestBoard_remaining(t *testing.T) {
	b := newBoard(3, 3)
	b.load(1)

	assert.Equal(t, 1, b.remaining())
	b.flag(0, 0)
	assert.Equal(t, 0, b.remaining())
}

func TestBoard_reveal(t *testing.T) {
	b := newBoard(3, 3)
	b.load(0)
	b.setMine(1, 1)

	b.reveal(2, 1)
	assert.True(t, b.bugs[1][2].shown)
}

func TestBoard_revealCascade(t *testing.T) {
	b := newBoard(3, 4)
	b.load(0)
	b.setMine(1, 1)

	b.reveal(3, 1)
	assert.True(t, b.bugs[0][3].shown)
	assert.True(t, b.bugs[1][3].shown)
	assert.True(t, b.bugs[2][3].shown)
}
