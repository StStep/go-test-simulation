package formation

import (
	pr "github.com/StStep/go-test-simulation/internal/formationprop"
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
		{"", 1, .25, .25, 1, 0, [2]float64{0, 0}},        // Single unit
	}

	for i, v := range tables {
		f := NewFormation(pr.NewFormationProp(v.style, v.width, v.fileSpacing, v.rankSpacing), v.size)
		assert.Equal(v.size, f.size, "Test %v", i)
		assert.Equal(v.expCapSlot, f.capSlot, "Test %v", i)
		assert.Equal("Captain", f.tags[f.capSlot], "Test %v", i)
		for k := 0; k < f.size; k++ {
			x := v.firstOffset[0] + float64(k%v.width)*v.fileSpacing
			y := v.firstOffset[1] - float64(k/v.width)*v.rankSpacing
			assert.InDeltaf(x, f.Offset(k)[0], 0.01, "Test %v, 0 Loop %v", i, k)
			assert.InDeltaf(y, f.Offset(k)[1], 0.01, "Test %v, 0 Loop %v", i, k)
		}
	}
}

func TestGiveTakeSlot(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		style       string
		width       int
		fileSpacing float64
		rankSpacing float64
		size        int
		notEmpty    []int
		give        int
		take        int
		expFound    bool
	}{
		{"", 5, .25, .25, 20, []int{0, 1, 2, 3}, 1, 1, true},  // Start full, give then take
		{"", 5, .25, .25, 20, []int{0, 1, 2, 3}, 1, 3, false}, // Start full, give then take
	}

	for i, v := range tables {
		f := NewFormation(pr.NewFormationProp(v.style, v.width, v.fileSpacing, v.rankSpacing), v.size)
		for _, closedSlot := range v.notEmpty {
			assert.Truef(f.TakeSlot(closedSlot), "Test %v", i)
		}
		f.GiveSlot(v.give)
		assert.Equalf(v.expFound, f.TakeSlot(v.take), "Test %v", i)
	}
}

func TestTakeClosestOpenSlot(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		style       string
		width       int
		fileSpacing float64
		rankSpacing float64
		size        int
		pos         [2]float64
		tag         string
		notEmpty    []int
		expSlot     int
		expFound    bool
	}{
		{"", 5, .25, .25, 20, [2]float64{0.0, 0.0}, "Captain", []int{0, 1, 2, 3}, 7, true}, // Empty Captain slot Basic block
		{"", 5, .25, .25, 20, [2]float64{0.0, 0.0}, "Captain", []int{7}, -1, false},        // Full Captain slot Basic block
		{"", 5, .25, .25, 20, [2]float64{-0.1, 0.25}, "Grunt", []int{3}, 2, true},          // Grunt slot Basic block
		{"", 5, .25, .25, 20, [2]float64{-0.1, 0.25}, "Grunt", []int{2}, 1, true},          // Full Grunt slot Basic block
	}

	for i, v := range tables {
		f := NewFormation(pr.NewFormationProp(v.style, v.width, v.fileSpacing, v.rankSpacing), v.size)
		for _, closedSlot := range v.notEmpty {
			assert.Truef(f.TakeSlot(closedSlot), "Test %v", i)
		}
		slot, isFound := f.TakeClosestOpenSlot(v.tag, v.pos)
		assert.Equalf(v.expSlot, slot, "Test %v", i)
		assert.Equalf(v.expFound, isFound, "Test %v", i)
	}
}
