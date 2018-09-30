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
			assert.Truef(s.Contains(k+1), "Test %v", i)
		}
		assert.Equalf(v.expCount, s.EntityCount(), "Test %v", i)

		for k := 0; k < v.count; k++ {
			s.UnregisterEntity(k + 1)
			assert.Falsef(s.Contains(k+1), "Test %v", i)
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
			assert.Falsef(s.UpdateEntity(v.id, v.vel), "Test %v", i)
		}

		for k := 0; k < v.count; k++ {
			s.UnregisterEntity(k + 1)
		}
		assert.Equalf(0, s.EntityCount(), "Test %v", i)
	}
}

func TestStep(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		startPos [2]float64
		vel      [2]float64
		stepSize float64
		steps    int
		expPos   [2]float64
	}{
		{[2]float64{0, 0}, [2]float64{10, 0}, 0.01, 100, [2]float64{10, 0}}, // Forward 10
	}

	for i, v := range tables {
		s := NewSpace()
		s.RegisterEntity(v.startPos, 0)
		s.UpdateEntity(1, v.vel)

		for k := 0; k < v.steps; k++ {
			s.Step(v.stepSize)
		}
		pos := s.positions[1]
		assert.InDeltaSlicef(v.expPos[:], pos[:], 0.01, "Test %v: Exp %v Pos %v", i, v.expPos, pos)
	}
}

func TestCollisions(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		poss    [][2]float64
		radii   []float64
		expColl [][2]int
	}{
		{[][2]float64{[2]float64{0, 0}, [2]float64{2, 0}}, []float64{1, 1.5},
			[][2]int{[2]int{1, 2}}}, // Simple Collision
		{[][2]float64{[2]float64{0, 0}, [2]float64{2, 0}, [2]float64{0, 1}}, []float64{1, 1.5, 2},
			[][2]int{[2]int{1, 2}, [2]int{1, 3}, [2]int{2, 3}}}, // Double Collision
		{[][2]float64{[2]float64{0, 0}, [2]float64{2, 0}}, []float64{1, 0.5}, [][2]int{}}, // No Collision
	}

	for i, v := range tables {
		s := NewSpace()

		for k := 0; k < len(v.poss); k++ {
			s.RegisterEntity(v.poss[k], v.radii[k])
		}
		coll := s.Collisions()

		assert.Equalf(len(v.expColl), len(coll), "Test %v", i)
		for k := 0; k < len(v.expColl); k++ {
			assert.Equalf(v.expColl[k][0], coll[k][0], "Test %v", i)
			assert.Equalf(v.expColl[k][1], coll[k][1], "Test %v", i)
		}
	}
}
