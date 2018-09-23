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
		{[2]float64{0, 1}, 4, [2]float64{0, 1}, 4},           // Forward Equal
		{[2]float64{1, 0}, 4, [2]float64{1, 0}, 3},           // Right Saturate Vel
		{[2]float64{0, 2}, 8, [2]float64{0, 1}, 4},           // Forward Saturate Dir and Vel
		{[2]float64{2, 0}, 8, [2]float64{1, 0}, 3},           // Right Saturate Dir and Vel
		{[2]float64{-1, 0}, 4, [2]float64{-1, 0}, 2},         // Left Saturate Vel
		{[2]float64{0, -1}, 4, [2]float64{0, -1}, 1},         // Backwards Saturate Vel
		{[2]float64{.5, .5}, 4, [2]float64{.5, .5}, 3.5},     // Forward-Right Saturate Vel
		{[2]float64{-.5, -.5}, 4, [2]float64{-.5, -.5}, 1.5}, // Back-Left Saturate Vel
		{[2]float64{0, 0.5}, 2, [2]float64{0, 1}, 2},         // Forward Less Dir and Ve
		{[2]float64{1, 0}, 1, [2]float64{1, 0}, 1},           // Right Less Vel
		{[2]float64{-1, 0}, 1, [2]float64{-1, 0}, 1},         // Left Less Vel
		{[2]float64{0, -1}, 0.5, [2]float64{0, -1}, 0.5},     // Backwards Less Vel
	}

	for _, v := range tables {
		m := NewMovement([4]float64{4, 1, 3, 2}, [4]float64{}, [4]float64{}, [4]float64{}, [4]float64{}, 0, 0)
		m.SetCommand(v.dir, v.cmdSpeed)
		assert.Equal(v.expDir, m.CmdDirection)
		assert.Equal(v.expSpeed, m.CmdSpeed)
	}
}

func TestTurnRate(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		baseRadius    float64
		rate          float64
		speed         float64
		expTurnRadius float64
	}{
		{3, 0.5, 2, 4},  // Normal
		{3, -0.5, 2, 2}, // Negative
		{3, -1, 3, 0},   // Exact 0
		{3, -1, 10, 0},  // Beyond 0
		{0, 1, 0, 0},    // 0 Base and No Vel
		{0, 1, 2, 2},    // 0 Base and Vel
	}

	for _, v := range tables {
		m := NewMovement([4]float64{}, [4]float64{}, [4]float64{}, [4]float64{}, [4]float64{}, v.baseRadius, v.rate)
		m.CurVelocity = [2]float64{0, v.speed}
		assert.Equal(v.expTurnRadius, m.Properties.TurnRateAt(v.speed))
		assert.Equal(v.expTurnRadius, m.TurnRate())
	}
}
