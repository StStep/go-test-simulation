package space

import (
	"github.com/StStep/go-test-simulation/internal/id"
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
			id := id.Eid(k + 1)
			s.RegisterEntity(id, v.pos, v.radius)
			assert.Truef(s.UpdateEntity(id, v.vel), "Test %v", i)
			assert.Truef(s.Contains(id), "Test %v", i)
		}
		assert.Equalf(v.expCount, s.EntityCount(), "Test %v", i)

		for k := 0; k < v.count; k++ {
			id := id.Eid(k + 1)
			s.UnregisterEntity(id)
			assert.Falsef(s.Contains(id), "Test %v", i)
		}
		assert.Equalf(0, s.EntityCount(), "Test %v", i)
	}
}

func TestUpdateEntity(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		count int
		id    id.Eid
		vel   [2]float64
		exp   bool
	}{
		{4, id.Eid(1), [2]float64{2, 5}, true},  // Standard usage
		{4, id.Eid(5), [2]float64{1, 2}, false}, // Unknown ID error
	}

	for i, v := range tables {
		s := NewSpace()
		for k := 0; k < v.count; k++ {
			s.RegisterEntity(id.Eid(k+1), [2]float64{}, 0)
		}
		if v.exp {
			assert.Truef(s.UpdateEntity(v.id, v.vel), "Test %v", i)
			vel := s.Velocity(v.id)
			assert.InDeltaSlicef(vel[:], v.vel[:], 0.01, "Test %v", i)
		} else {
			assert.Falsef(s.UpdateEntity(v.id, v.vel), "Test %v", i)
		}

		for k := 0; k < v.count; k++ {
			s.UnregisterEntity(id.Eid(k + 1))
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
		id := id.Eid(1)
		s.RegisterEntity(id, v.startPos, 0)
		s.UpdateEntity(id, v.vel)

		for k := 0; k < v.steps; k++ {
			s.Step(v.stepSize)
		}
		pos := s.Position(id)
		assert.InDeltaSlicef(v.expPos[:], pos[:], 0.01, "Test %v: Exp %v Pos %v", i, v.expPos, pos)
	}
}

func TestCollisions(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		poss    [][2]float64
		radii   []float64
		expColl [][2]id.Eid
	}{
		{[][2]float64{[2]float64{0, 0}, [2]float64{2, 0}}, []float64{1, 1.5},
			[][2]id.Eid{[2]id.Eid{1, 2}}}, // Simple Collision
		{[][2]float64{[2]float64{0, 0}, [2]float64{2, 0}, [2]float64{0, 1}}, []float64{1, 1.5, 2},
			[][2]id.Eid{[2]id.Eid{1, 2}, [2]id.Eid{1, 3}, [2]id.Eid{2, 3}}}, // Double Collision
		{[][2]float64{[2]float64{0, 0}, [2]float64{2, 0}}, []float64{1, 0.5}, [][2]id.Eid{}}, // No Collision
	}

	for i, v := range tables {
		s := NewSpace()

		for k := 0; k < len(v.poss); k++ {
			s.RegisterEntity(id.Eid(k+1), v.poss[k], v.radii[k])
		}
		coll := s.Collisions()

		assert.Equalf(len(v.expColl), len(coll), "Test %v", i)
		for k := 0; k < len(v.expColl); k++ {
			assert.Equalf(v.expColl[k][0], coll[k][0], "Test %v", i)
			assert.Equalf(v.expColl[k][1], coll[k][1], "Test %v", i)
		}
	}
}
