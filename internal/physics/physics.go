package physics

// TODO Consider optimizing, see https://cstheory.stackexchange.com/questions/16927/efficient-algorithm-to-find-overlapping-circles-of-various-sizes
// Note: Leaves control of velocity in hands of updater, only provides collision information, doesn't enforce it

import (
	"fmt"
	"github.com/StStep/go-test-simulation/internal/physics/inertia"
	pr "github.com/StStep/go-test-simulation/internal/physics/prop"
	"github.com/go-logfmt/logfmt"
	fl "gonum.org/v1/gonum/floats"
	"io"
)

type Physics struct {
	ids       []uint64
	inertia   map[uint64]*inertia.Inertia
	positions map[uint64][2]float64
	radii     map[uint64]float64
	cvalid    bool
	ccolls    [][2]uint64
	loge      *logfmt.Encoder
}

func New() *Physics {
	var s Physics
	s.ids = make([]uint64, 0)
	s.inertia = make(map[uint64]*inertia.Inertia)
	s.positions = make(map[uint64][2]float64)
	s.radii = make(map[uint64]float64)
	s.cvalid = false
	s.ccolls = make([][2]uint64, 0)
	s.loge = nil

	return &s
}

func (s *Physics) SetLogOutput(out io.Writer) {
	s.loge = logfmt.NewEncoder(out)
}

func (s *Physics) EntityCount() int {
	return len(s.positions)
}

func (s *Physics) Contains(id uint64) bool {
	_, ok1 := s.positions[id]
	_, ok2 := s.radii[id]
	_, ok3 := s.inertia[id]
	ok4 := false
	for _, v := range s.ids {
		if v == id {
			ok4 = true
			break
		}
	}
	return ok1 && ok2 && ok3 && ok4
}

func (s *Physics) RegisterEntity(id uint64, prop *pr.Prop, pos [2]float64) {
	s.positions[id] = pos
	s.radii[id] = prop.FootprintRadius
	s.inertia[id] = inertia.NewInertia(prop)
	s.ids = append(s.ids, id)
	if s.loge != nil {
		s.loge.EncodeKeyval("tag", "add")
		s.loge.EncodeKeyval("id", id)
		s.loge.EncodeKeyval("shape", "circle")
		s.loge.EncodeKeyval("pos", fmt.Sprintf("%v,%v", pos[0], pos[1]))
		s.loge.EncodeKeyval("radius", s.radii[id])
		s.loge.EndRecord()
	}
}

func (s *Physics) Command(id uint64) (dir [2]float64, speed float64) {
	return s.inertia[id].Command()
}

func (s *Physics) SetCommand(id uint64, dir [2]float64, speed float64) {
	s.inertia[id].SetCommand(dir, speed)
}

func (s *Physics) UnregisterEntity(id uint64) {
	// Delete id entries from maps
	delete(s.positions, id)
	delete(s.radii, id)
	delete(s.inertia, id)

	// Remove id from list
	for i, v := range s.ids {
		if v == id {
			s.ids = append(s.ids[:i], s.ids[i+1:]...)
			break
		}
	}
}

func (s *Physics) Step(del float64) {

	if s.loge != nil {
		s.loge.EncodeKeyval("tag", "step")
		s.loge.EndRecord()
	}

	// Calc for each
	for _, v := range s.ids {
		s.inertia[v].PhyStep(del)
		pos := s.positions[v]
		vel := s.Velocity(v)
		svel := vel[:]
		fl.Scale(del, svel)
		fl.Add(pos[:], svel)
		var res [2]float64
		copy(res[:], pos[:])
		s.positions[v] = res

		if s.loge != nil {
			s.loge.EncodeKeyval("tag", "update")
			s.loge.EncodeKeyval("id", v)
			s.loge.EncodeKeyval("pos", fmt.Sprintf("%v,%v", pos[0], pos[1]))
			cmddir, cmdspeed := s.Command(v)
			s.loge.EncodeKeyval("cmddir", fmt.Sprintf("%v,%v", cmddir[0], cmddir[1]))
			s.loge.EncodeKeyval("cmdspeed", cmdspeed)
			s.loge.EncodeKeyval("vel", fmt.Sprintf("%v,%v", vel[0], vel[1]))
			s.loge.EndRecord()
		}
	}

	// Invaidate and empty
	s.cvalid = false
	s.ccolls = s.ccolls[:0]
}

func (s *Physics) Position(id uint64) [2]float64 {
	return s.positions[id]
}

func (s *Physics) Velocity(id uint64) [2]float64 {
	return s.inertia[id].Velocity()
}

func (s *Physics) Collisions() [][2]uint64 {
	// Send cached if valid
	if s.cvalid {
		return s.ccolls
	}

	// Otherwise calc
	for i, v := range s.ids {
		pos := s.positions[v]
		rad := s.radii[v]
		for k := i + 1; k < len(s.ids); k++ {
			tid := s.ids[k]
			tpos := s.positions[tid]
			if fl.Distance(pos[:], tpos[:], 2) < rad+s.radii[tid] {
				s.ccolls = append(s.ccolls, [2]uint64{v, tid})
			}
		}
	}
	s.cvalid = true
	return s.ccolls
}

// TODO Return collision list filtered for specific ID
func (s *Physics) EntityCollisions(id int) []int {
	return []int{}
}

// TODO Consider the following checks:
// 1. Compare currnet dist, radii and movement distance to see if collision is even possible
// 2. Comparing relative velocity and current distance to know if conflict could happen
// 3. Finally do detailed check to see if overlapping will occur?
func (s *Physics) ProjectedCollisions() [][2]int {
	return [][2]int{}
}

// TODO Return projected collision list filtered for specific ID
func (s *Physics) ProjectedEntityCollisions(id int) []int {
	return []int{}
}
