package unit

type Position struct {
	X, Y float64
}

type Member struct {
	Position Position
}

type Unit struct {
	Members []Member
}

func NewUnit(size int, pos Position) *Unit {
	if size <= 0 {
		return nil
	}
	members := make([]Member, size)
	for i := 0; i < size; i++ {
		members[i].Position = pos
	}
	return &Unit{members}
}

func (u *Unit) GetSize() int {
	return len(u.Members)
}

func (u *Unit) GetPosition() Position {
	var totalX, totalY, len float64 = 0, 0, float64(len(u.Members))
	for _, v := range u.Members {
		totalX += v.Position.X
		totalY += v.Position.Y
	}
	return Position{X: totalX / len, Y: totalY / len}
}

func (u *Unit) UpdateMove(done chan bool) {
	done <- false
}
