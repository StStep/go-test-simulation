package formation

import (
	"fmt"
	fl "gonum.org/v1/gonum/floats"
	"strings"
)

// TODO Purpose of Formation
//
// Formation does the math derived from FormationProp and given info (size)
// and provides interface for entities to be assign slots in it
//
// * Leader is important slots, all others reference it
// * Want to work with offsets
// * Consumed by Unit, members only know of assigned offset and orientation to leader
// * Want to be able to trade slots, allow for the weak/tired to move to the back/center
// * Need to be able to change FormationProps somehow, allow for reinitialization?
// * Following logic is built around standard block formation with leader in center
// * Need to allow for varying formation styles, such as square formation

// Rmove capSlot junk, using guide now
type Formation struct {
	Prop    *Prop
	size    int
	capSlot int
	slots   [][2]float64
	tags    []string
	isEmpty []bool
}

type Prop struct {
	Style       string  `json:"style"`
	Width       int     `json:"width"`
	FileSpacing float64 `json:"fileSpacing"`
	RankSpacing float64 `json:"rankSpacing"`
}

func NewFormation(prop *Prop, size int) *Formation {
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
	if f.Prop.Style == "" {
		// Assuming standard block with captain in middle
		ranks := f.size / f.Prop.Width
		hwidth := f.Prop.Width - f.Prop.Width/2 - 1
		hrank := ranks - ranks/2 - 1
		f.capSlot = f.Prop.Width*(hrank) + hwidth

		for i := 0; i < f.size; i++ {
			if i == f.capSlot {
				f.tags[i] = "Captain"
				f.slots[i] = [2]float64{0.0, 0.0}
			} else {
				f.tags[i] = "Grunt"
				fInd := i % f.Prop.Width
				rInd := i / f.Prop.Width
				f.slots[i] = [2]float64{float64(fInd-hwidth) * f.Prop.FileSpacing, float64(hrank-rInd) * f.Prop.RankSpacing}
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
	rows := make([]string, 0, f.size+f.size/f.Prop.Width) // one per unit and per end-of-rank for \n
	for i, v := range f.slots {
		rows = append(rows, fmt.Sprint(v)+"\t")
		// Append newline at end-of-ranks
		if i%f.Prop.Width == f.Prop.Width-1 {
			rows = append(rows, "\n")
		}
	}
	return strings.Join(rows, "")
}
