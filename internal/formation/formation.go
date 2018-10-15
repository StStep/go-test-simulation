package formation

import (
	"fmt"
	pr "github.com/StStep/go-test-simulation/internal/formationprop"
	fl "gonum.org/v1/gonum/floats"
	"strings"
)

// TODO Formation should be more generic?
// How does it inter-relate with unit
// Should slots still be a thing?

type Formation struct {
	Prop    pr.FormationProp
	size    int
	capSlot int
	slots   [][2]float64
	tags    []string
	isEmpty []bool
}

func NewFormation(prop pr.FormationProp, size int) *Formation {
	f := Formation{Prop: prop}
	f.Resize(size)
	return &f
}

func (f *Formation) Resize(size int) {
	// Change slices first
	if f.size == 0 {
		f.slots = make([][2]float64, size)
		f.tags = make([]string, size)
		f.isEmpty = make([]bool, size)
	} else {
		panic("Not implemented") // TODO
	}
	f.size = size

	// Recalc positions
	if f.Prop.Style() == "" {
		// Assuming standard block with captain in middle
		ranks := f.size / f.Prop.Width()
		hwidth := f.Prop.Width() - f.Prop.Width()/2 - 1
		hrank := ranks - ranks/2 - 1
		f.capSlot = f.Prop.Width()*(hrank) + hwidth

		for i := 0; i < f.size; i++ {
			if i == f.capSlot {
				f.tags[i] = "Captain"
				f.slots[i] = [2]float64{0.0, 0.0}
			} else {
				f.tags[i] = "Grunt"
				fInd := i % f.Prop.Width()
				rInd := i / f.Prop.Width()
				f.slots[i] = [2]float64{float64(fInd-hwidth) * f.Prop.FileSpacing(), float64(hrank-rInd) * f.Prop.RankSpacing()}
			}
			f.isEmpty[i] = true
		}
	} else {
		panic("Not implemented") // TODO
	}
}

func (f *Formation) TakeClosestOpenSlot(tag string, rpos [2]float64) (int, bool) {
	minSlot := -1
	minDist := 0.0
	for i := 0; i < f.size; i++ {
		if f.isEmpty[i] && f.tags[i] == tag {
			slot := f.slots[i]
			dist := fl.Distance(rpos[:], slot[:], 2)
			if minSlot == -1 || dist < minDist {
				minSlot = i
				minDist = dist
			}
		}
	}
	return minSlot, minSlot != -1
}

func (f *Formation) TakeSlot(i int) bool {
	if f.isEmpty[i] {
		f.isEmpty[i] = false
		return true
	} else {
		return false
	}
}

func (f *Formation) GiveSlot(i int) {
	f.isEmpty[i] = true
}

func (f *Formation) Offset(i int) [2]float64 {
	return f.slots[i]
}

func (f *Formation) DebugOffsets() string {
	rows := make([]string, 0, f.size+f.size/f.Prop.Width()) // one per unit and per end-of-rank for \n
	for i, v := range f.slots {
		rows = append(rows, fmt.Sprint(v)+"\t")
		// Append newline at end-of-ranks
		if i%f.Prop.Width() == f.Prop.Width()-1 {
			rows = append(rows, "\n")
		}
	}
	return strings.Join(rows, "")
}
