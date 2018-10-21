package inertia

import (
	pr "github.com/StStep/go-test-simulation/internal/physics/prop"
	fl "gonum.org/v1/gonum/floats"
)

type Inertia struct {
	Prop        pr.Prop    // Physics properties to use with math
	curVelocity [2]float64 // Represents current velocity vector
	cmdVelocity [2]float64 // Represents commanded velocity vector
}

// Forward, Backward, Right, Left
func NewInertia(prop pr.Prop) *Inertia {
	return &Inertia{Prop: prop}
}

// Turn rate used for setting arc for current direction
func (m *Inertia) TurnRate() float64 {
	return m.Prop.TurnRateAt(fl.Norm(m.curVelocity[:], 2))
}

func (m *Inertia) Velocity() [2]float64 {
	return m.cmdVelocity
}

func (m *Inertia) Command() ([2]float64, float64) {
	t := m.cmdVelocity[:]
	speed := fl.Norm(t, 2)
	fl.Scale(1/speed, t)

	return [2]float64{t[0], t[1]}, speed
}

// dir[1] > 0 ? Front : Back; dir[0] > 0 ? Right : Left
func (m *Inertia) SetCommand(dir [2]float64, speed float64) {
	// Set to unit vector if not already
	fl.Scale(1/fl.Norm(dir[:], 2), dir[:])

	// Check max horizantal velocity
	hsp := dir[0]
	if hsp > 0 {
		hsp *= m.Prop.MaxVelocity()[2]
	} else {
		hsp *= m.Prop.MaxVelocity()[3]
	}

	// Check max vertical max velocity
	vsp := dir[1]
	if vsp > 0 {
		vsp *= m.Prop.MaxVelocity()[0]
	} else {
		vsp *= m.Prop.MaxVelocity()[1]
	}

	// Cap based on calc max
	adjSpeed := speed
	if mx := fl.Norm([]float64{hsp, vsp}, 2); adjSpeed > mx {
		adjSpeed = mx
	}

	// Set vel to dir scaled by speed
	copy(m.cmdVelocity[:], dir[:])
	fl.Scale(adjSpeed, m.cmdVelocity[:])
}

func (m *Inertia) PhyStep(del float64) {
	var diff [2]float64
	fl.SubTo(diff[:], m.curVelocity[:], m.cmdVelocity[:])

	// Check horizantal velocity, use right if positive
	hind := 3
	if m.curVelocity[0] > 0 {
		hind = 2
	} else if m.curVelocity[0] < 0 {
		hind = 3
	} else if m.cmdVelocity[0] > 0 {
		hind = 2
	} else {
		hind = 3
	}

	hdiff := diff[0]
	if hdiff > 0 {
		hdiff -= m.Prop.Deceleration()[hind] * del
		if hdiff < 0 {
			hdiff = 0
		}
	} else if hdiff < 0 {
		hdiff += m.Prop.Acceleration()[hind] * del
		if hdiff > 0 {
			hdiff = 0
		}
	} else {
		hdiff = 0
	}

	// Check vertival velocity, use forward if positive
	vind := 1
	if m.curVelocity[1] > 0 {
		vind = 0
	} else if m.curVelocity[1] < 0 {
		vind = 1
	} else if m.cmdVelocity[1] > 0 {
		vind = 0
	} else {
		vind = 1
	}

	vdiff := diff[1]
	if vdiff > 0 {
		vdiff -= m.Prop.Deceleration()[vind] * del
		if vdiff < 0 {
			vdiff = 0
		}
	} else if vdiff < 0 {
		vdiff += m.Prop.Acceleration()[vind] * del
		if vdiff > 0 {
			vdiff = 0
		}
	} else {
		vdiff = 0
	}

	// PhyStep  vel
	fl.AddTo(m.curVelocity[:], []float64{hdiff, vdiff}, m.cmdVelocity[:])
}
