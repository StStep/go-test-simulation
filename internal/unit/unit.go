package unit

type Member struct {
	Position [2]float64
}

type Unit struct {
	Members []Member
}

func NewUnit(size int, pos [2]float64) *Unit {
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

func (u *Unit) Position() [2]float64 {
	var totalX, totalY, len float64 = 0, 0, float64(len(u.Members))
	for _, v := range u.Members {
		totalX += v.Position[0]
		totalY += v.Position[1]
	}
	return [2]float64{totalX / len, totalY / len}
}

func (u *Unit) UpdateMove(done chan bool) {
	done <- false
}
