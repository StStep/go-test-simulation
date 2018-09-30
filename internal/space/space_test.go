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
			assert.Equalf(k+1, s.RegisterEntity(v.pos, v.radius), "Test %v", i)
			assert.Truef(s.UpdateEntity(k+1, v.vel), "Test %v", i)
		}
		assert.Equalf(v.expCount, s.EntityCount(), "Test %v", i)

		for k := 0; k < v.count; k++ {
			s.UnregisterEntity(k + 1)
		}
		assert.Equalf(0, s.EntityCount(), "Test %v", i)
	}
}

func TestUpdateEntity(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		count int
		id    int
		vel   [2]float64
		exp   bool
	}{
		{4, 1, [2]float64{2, 5}, true},  // Standard usage
		{4, 5, [2]float64{1, 2}, false}, // Unknown ID error
	}

	for i, v := range tables {
		s := NewSpace()
		for k := 0; k < v.count; k++ {
			assert.Equalf(k+1, s.RegisterEntity([2]float64{}, 0), "Test %v", i)
		}
		if v.exp {
			assert.Truef(s.UpdateEntity(v.id, v.vel), "Test %v", i)
			vel := s.velocity[v.id]
			assert.InDeltaSlicef(vel[:], v.vel[:], 0.01, "Test %v", i)
		} else {
			assert.False(s.UpdateEntity(v.id, v.vel), "Test %v", i)
		}

		for k := 0; k < v.count; k++ {
			s.UnregisterEntity(k + 1)
		}
		assert.Equal(0, s.EntityCount(), "Test %v", i)
	}
}
