package formation

import ()

type FormationProp struct {
	style       string
	width       int
	fileSpacing float64
	rankSpacing float64
}

type Formation struct {
	Prop    *FormationProp
	size    int
	capSlot int
	slots   [][2]float64
	tags    []string
	isEmpty []bool
}

func NewFormation(style string, width int, fileSpacing float64, rankSpacing float64, size int) *Formation {
	p := FormationProp{style, width, fileSpacing, rankSpacing}
	f := Formation{Prop: &p}
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
	if f.Prop.style == "" {
		// Assuming standard block with captain in middle
		ranks := f.size / f.Prop.width
		hwidth := f.Prop.width - f.Prop.width/2 - 1
		hrank := ranks - ranks/2 - 1
		f.capSlot = f.Prop.width*(hrank) + hwidth

		for i := 0; i < f.size; i++ {
			if i == f.capSlot {
				f.tags[i] = "Captain"
				f.slots[i] = [2]float64{0.0, 0.0}
			} else {
				f.tags[i] = "grunt"
				fInd := i % f.Prop.width
				rInd := i / f.Prop.width
				f.slots[i] = [2]float64{float64(fInd-hwidth) * f.Prop.fileSpacing, float64(hrank-rInd) * f.Prop.rankSpacing}
			}
			f.isEmpty[i] = false
		}
	} else {
		panic("Not implemented") // TODO
	}
}

func (f *Formation) ClosestOpenSlot(tag string, rpos [2]float64) int {
	return -1
}
