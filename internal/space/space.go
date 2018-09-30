package space

// TODO Consider optimizing, see https://cstheory.stackexchange.com/questions/16927/efficient-algorithm-to-find-overlapping-circles-of-various-sizes
// Note: Leaves control of velocity in hands of updater, only provides collision information, doesn't enforce it

import (
	fl "gonum.org/v1/gonum/floats"
)

type Space struct {
	positions map[int][2]float64
	radii     map[int]float64
	velocity  map[int][2]float64
	ids       []int
	lastId    int
	cvalid    bool
	ccolls    [][2]int
}

func NewSpace() *Space {
	var s Space
	s.positions = make(map[int][2]float64)
	s.radii = make(map[int]float64)
	s.velocity = make(map[int][2]float64)
	s.ids = make([]int, 0)
	s.cvalid = false
	s.ccolls = make([][2]int, 0)

	return &s
}

func (s *Space) EntityCount() int {
	return len(s.positions)
}

func (s *Space) Contains(id int) bool {
	_, ok1 := s.positions[id]
	_, ok2 := s.radii[id]
	_, ok3 := s.velocity[id]
	ok4 := false
	for _, v := range s.ids {
		if v == id {
			ok4 = true
			break
		}
	}
	return ok1 && ok2 && ok3 && ok4
}

func (s *Space) RegisterEntity(pos [2]float64, radius float64) int {
	s.lastId += 1
	s.positions[s.lastId] = pos
	s.radii[s.lastId] = radius
	s.velocity[s.lastId] = [2]float64{}
	s.ids = append(s.ids, s.lastId)
	return s.lastId
}

// TODO Invalidates prev collision check
func (s *Space) UpdateEntity(id int, vel [2]float64) bool {
	_, ok := s.velocity[id]
	if !ok {
		return false
	}
	s.velocity[id] = vel
	return true
}

func (s *Space) UnregisterEntity(id int) {
	// Delete id entries from maps
	delete(s.positions, id)
	delete(s.radii, id)
	delete(s.velocity, id)

	// Remove id from list
	for i, v := range s.ids {
		if v == id {
			s.ids = append(s.ids[:i], s.ids[i+1:]...)
			break
		}
	}
}

func (s *Space) Step(del float64) {

	// Calc for each
	for _, v := range s.ids {
		pos := s.positions[v]
		vel := s.velocity[v]
		svel := vel[:]
		fl.Scale(del, svel)
		fl.Add(pos[:], svel)
		var res [2]float64
		copy(res[:], pos[:])
		s.positions[v] = res
	}

	// Invaidate and empty
	s.cvalid = false
	s.ccolls = s.ccolls[:0]
}

func (s *Space) Collisions() [][2]int {
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
				s.ccolls = append(s.ccolls, [2]int{v, tid})
			}
		}
	}
	s.cvalid = true
	return s.ccolls
}

// TODO Return collision list filtered for specific ID
func (s *Space) EntityCollisions(id int) []int {
	return []int{}
}

// TODO Consider the following checks:
// 1. Compare currnet dist, radii and movement distance to see if collision is even possible
// 2. Comparing relative velocity and current distance to know if conflict could happen
// 3. Finally do detailed check to see if overlapping will occur?
func (s *Space) ProjectedCollisions() [][2]int {
	return [][2]int{}
}

// TODO Return projected collision list filtered for specific ID
func (s *Space) ProjectedEntityCollisions(id int) []int {
	return []int{}
}
