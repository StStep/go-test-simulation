package space

// TODO Consider optimizing, see https://cstheory.stackexchange.com/questions/16927/efficient-algorithm-to-find-overlapping-circles-of-various-sizes

type Space struct {
	positions map[int][2]float64
	radii     map[int]float64
	velocity  map[int][2]float64
	lastId    int
}

func NewSpace() *Space {
	var s Space
	s.positions = make(map[int][2]float64)
	s.radii = make(map[int]float64)
	s.velocity = make(map[int][2]float64)

	return &s
}

func (s *Space) EntityCount() int {
	return len(s.positions)
}

func (s *Space) RegisterEntity(pos [2]float64, radius float64) int {
	s.lastId += 1
	s.positions[s.lastId] = pos
	s.radii[s.lastId] = radius
	s.velocity[s.lastId] = [2]float64{}
	return s.lastId
}

func (s *Space) UpdateEntity(id int, vel [2]float64) bool {
	_, ok := s.velocity[id]
	if !ok {
		return false
	}
	s.velocity[id] = vel
	return true
}

func (s *Space) UnregisterEntity(id int) {
	delete(s.positions, id)
	delete(s.radii, id)
	delete(s.velocity, id)
}

func (s *Space) Step(del float64) {

}

func (s *Space) Collisions() [][2]int {
	return [][2]int{}
}

func (s *Space) EntityCollisions(id int) []int {
	return []int{}
}

func (s *Space) ProjectedCollisions() [][2]int {
	return [][2]int{}
}

func (s *Space) ProjectedEntityCollisions(id int) []int {
	return []int{}
}
