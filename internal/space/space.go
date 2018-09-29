package space

type Space struct {
	positions [][2]float64
	radii     []float64
	veclocity [][2]float64
}

func (s *Space) RegisterEntity(pos [2]float64, radius float64) int {
	return -1
}

func (s *Space) UpdateEntity(i int, vel float64) {
}

func (s *Space) UnregisterEntity(i int) {
}

func (s *Space) Step() {

}

func (s *Space) Collisions() [][2]float64 {
	return [][2]float64{}
}

func (s *Space) EntityCollisions(i int) []float64 {
	return []float64{}
}

func (s *Space) ProjectedCollisions() [][2]float64 {
	return [][2]float64{}
}

func (s *Space) ProjectedEntityCollisions(i int) []float64 {
	return []float64{}
}
