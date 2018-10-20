package space

import (
	"github.com/StStep/go-test-simulation/internal/id"
	pr "github.com/StStep/go-test-simulation/internal/physics/prop"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisteration(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		count    int
		expCount int
	}{
		{2, 2},   // Add 2
		{12, 12}, // Add 12
	}

	for i, v := range tables {
		prop := pr.NewProp([4]float64{}, [4]float64{}, [4]float64{}, [4]float64{}, [4]float64{}, 0, 0, 0)
		s := NewSpace()
		for k := 0; k < v.count; k++ {
			id := id.Eid(k + 1)
			s.RegisterEntity(id, prop, [2]float64{0, 0})
			s.UpdateEntity(id, [2]float64{0, 0}, 0)
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

func TestStep(t *testing.T) {
	assert := assert.New(t)

	tables := []struct {
		startPos [2]float64
		dir      [2]float64
		speed    float64
		stepSize float64
		steps    int
		expVel   [2]float64
		expPos   [2]float64
	}{
		{[2]float64{0, 0}, [2]float64{1, 0}, 10, 0.01, 100, [2]float64{10, 0}, [2]float64{10, 0}}, // Forward 10
	}

	for i, v := range tables {
		prop := pr.NewProp([4]float64{10, 10, 10, 10}, [4]float64{10, 10, 10, 10}, [4]float64{}, [4]float64{}, [4]float64{}, 0, 0, 0)
		s := NewSpace()
		id := id.Eid(1)
		s.RegisterEntity(id, prop, v.startPos)
		s.UpdateEntity(id, v.dir, v.speed)

		for k := 0; k < v.steps; k++ {
			s.Step(v.stepSize)
		}
		pos := s.Position(id)
		vel := s.Velocity(id)
		assert.InDeltaSlicef(v.expVel[:], vel[:], 0.01, "Test %v: Exp %v Vel %v", i, v.expVel, vel)
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
			prop := pr.NewProp([4]float64{}, [4]float64{}, [4]float64{}, [4]float64{}, [4]float64{}, 0, 0, v.radii[k])
			s.RegisterEntity(id.Eid(k+1), prop, v.poss[k])
		}
		coll := s.Collisions()

		assert.Equalf(len(v.expColl), len(coll), "Test %v", i)
		for k := 0; k < len(v.expColl); k++ {
			assert.Equalf(v.expColl[k][0], coll[k][0], "Test %v", i)
			assert.Equalf(v.expColl[k][1], coll[k][1], "Test %v", i)
		}
	}
}
