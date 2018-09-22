package movement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetCommand(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		dir      [2]float64
		cmdSpeed float64
		expDir   [2]float64
		expSpeed float64
	}{
		{[2]float64{0, 1}, 4, [2]float64{0, 1}, 4},
		{[2]float64{1, 0}, 4, [2]float64{1, 0}, 3},
		{[2]float64{0, 2}, 8, [2]float64{0, 1}, 4},
		{[2]float64{2, 0}, 8, [2]float64{1, 0}, 3},
		{[2]float64{-1, 0}, 4, [2]float64{-1, 0}, 2},
		{[2]float64{0, -1}, 4, [2]float64{0, -1}, 1},
		{[2]float64{.5, .5}, 4, [2]float64{.5, .5}, 3.5},
		{[2]float64{-.5, -.5}, 4, [2]float64{-.5, -.5}, 1.5},
	}

	for _, v := range tables {
		m := NewMovement([4]float64{4, 1, 3, 2})
		m.SetCommand(v.dir, v.cmdSpeed)
		assert.Equal(v.expDir, m.CmdDirection)
		assert.Equal(v.expSpeed, m.CmdSpeed)
	}
}
