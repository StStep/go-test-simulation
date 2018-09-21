package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnit(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		size      int
		pos       [2]float64
		expectNil bool
	}{
		{1, [2]float64{0, 0}, false},
		{1, [2]float64{4, 5}, false},
		{10, [2]float64{0, 0}, false},
		{10, [2]float64{-5, -4}, false},
		{0, [2]float64{0, 0}, true},
		{-5, [2]float64{0, 0}, true},
	}

	for _, v := range tables {
		u := NewUnit(v.size, v.pos)
		if v.expectNil {
			assert.Nil(u)
		} else {
			assert.NotNil(u)
		}
	}
}

func TestSize(t *testing.T) {
	assert := assert.New(t)

	sizes := []int{1, 4, 8, 20, 200}

	for _, v := range sizes {
		u := NewUnit(v, [2]float64{0, 0})
		assert.Equal(v, u.Size())
	}
}

func TestPosition(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		size int
		pos  [2]float64
	}{
		{1, [2]float64{0, 0}},
		{1, [2]float64{4, 5}},
		{10, [2]float64{}},
		{10, [2]float64{-5, -4}},
	}

	for _, v := range tables {
		u := NewUnit(v.size, v.pos)
		assert.Equal(v.pos, u.Position())
	}
}
