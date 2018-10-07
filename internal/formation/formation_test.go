package formation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFormation(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		style       string
		width       int
		fileSpacing float64
		rankSpacing float64
		size        int
		expCapSlot  int
		firstOffset [2]float64
	}{
		{"", 5, .25, .25, 20, 7, [2]float64{-0.5, 0.25}}, // Basic 5 file 4 rank block
	}

	for i, v := range tables {
		f := NewFormation(v.style, v.width, v.fileSpacing, v.rankSpacing, v.size)
		assert.Equal(v.size, f.size, "Test %v", i)
		assert.Equal(v.expCapSlot, f.capSlot, "Test %v", i)
		assert.Equal("Captain", f.tags[f.capSlot], "Test %v", i)
		for k := 0; k < f.size; k++ {
			assert.InDeltaf(v.firstOffset[0]+float64(k%v.width)*v.fileSpacing, f.slots[k][0], 0.01, "Test %v, 0 Loop %v", i, k)
			assert.InDeltaf(v.firstOffset[1]-float64(k/v.width)*v.rankSpacing, f.slots[k][1], 0.01, "Test %v, 1 Loop %v", i, k)
		}
	}
}
