package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnit(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		size      int
		pos       Position
		expectNil bool
	}{
		{1, Position{}, false},
		{1, Position{4, 5}, false},
		{10, Position{}, false},
		{10, Position{-5, -4}, false},
		{0, Position{}, true},
		{-5, Position{}, true},
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
		u := NewUnit(v, Position{})
		assert.Equal(v, u.Size())
	}
}

func TestPosition(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		size int
		pos  Position
	}{
		{1, Position{}},
		{1, Position{4, 5}},
		{10, Position{}},
		{10, Position{-5, -4}},
	}

	for _, v := range tables {
		u := NewUnit(v.size, v.pos)
		assert.Equal(v.pos, u.Position())
	}
}
