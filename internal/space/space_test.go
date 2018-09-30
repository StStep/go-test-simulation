package space

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisteration(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		pos      [2]float64
		radius   float64
		vel      [2]float64
		count    int
		expCount int
	}{
		{[2]float64{0, 0}, 4, [2]float64{0, 1}, 2, 2},   // Add 2
		{[2]float64{0, 0}, 4, [2]float64{0, 1}, 12, 12}, // Add 12
	}

	for i, v := range tables {
		s := NewSpace()
		for k := 0; k < v.count; k++ {
			assert.Equal(k+1, s.RegisterEntity(v.pos, v.radius), "Test %v", i)
			assert.True(s.UpdateEntity(k+1, v.vel))
		}
		assert.Equal(v.expCount, s.EntityCount(), "Test %v", i)

		for k := 0; k < v.count; k++ {
			s.UnregisterEntity(k + 1)
		}
		assert.Equal(0, s.EntityCount(), "Test %v", i)
	}
}
