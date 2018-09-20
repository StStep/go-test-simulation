package unit

import (
	"github.com/StStep/go-test-simulation/internal/vecmath"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnit(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		size      int
		pos       vecmath.Vector
		expectNil bool
	}{
		{1, vecmath.Vector{}, false},
		{1, vecmath.Vector{4, 5}, false},
		{10, vecmath.Vector{}, false},
		{10, vecmath.Vector{-5, -4}, false},
		{0, vecmath.Vector{}, true},
		{-5, vecmath.Vector{}, true},
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
		u := NewUnit(v, vecmath.Vector{})
		assert.Equal(v, u.Size())
	}
}

func TestPosition(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		size int
		pos  vecmath.Vector
	}{
		{1, vecmath.Vector{}},
		{1, vecmath.Vector{4, 5}},
		{10, vecmath.Vector{}},
		{10, vecmath.Vector{-5, -4}},
	}

	for _, v := range tables {
		u := NewUnit(v.size, v.pos)
		assert.Equal(v.pos, u.Position())
	}
}
