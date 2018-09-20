package unit

import (
	"github.com/StStep/go-test-simulation/internal/vecmath"
)

type Member struct {
	Position vecmath.Vector
}

type Unit struct {
	Members []Member
}

func NewUnit(size int, pos vecmath.Vector) *Unit {
	if size <= 0 {
		return nil
	}
	members := make([]Member, size)
	for i := 0; i < size; i++ {
		members[i].Position = pos
	}
	return &Unit{members}
}

func (u *Unit) Size() int {
	return len(u.Members)
}

func (u *Unit) Position() vecmath.Vector {
	var totalX, totalY, len float64 = 0, 0, float64(len(u.Members))
	for _, v := range u.Members {
		totalX += v.Position.X
		totalY += v.Position.Y
	}
	return vecmath.Vector{X: totalX / len, Y: totalY / len}
}

func (u *Unit) UpdateMove(done chan bool) {
	done <- false
}
