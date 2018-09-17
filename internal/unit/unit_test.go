package unit

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewUnit(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		size int
		pos Position
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

func TestGetSize(t *testing.T) {
}

func TestGetPosition(t *testing.T) {

}

