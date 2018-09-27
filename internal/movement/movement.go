package movement

import (
	fl "gonum.org/v1/gonum/floats"
)

type MoveProp struct {
	MaxVelocity                      [4]float64
	Acceleration, Deceleration       [4]float64
	EnergyUsageRate, BaseEnergyUsage [4]float64
	TurnRadiusRate, BaseTurnRadius   float64
}

func (m *MoveProp) TurnRateAt(speed float64) float64 {
	ret := m.BaseTurnRadius + m.TurnRadiusRate*speed
	if ret < 0 {
		return 0
	} else {
		return ret
	}
}

type Movement struct {
	Prop        *MoveProp  // Movement properties to use with math
	curVelocity [2]float64 // Represents current velocity vector
	cmdVelocity [2]float64 // Represents commanded velocity vector
}

// Forward, Backward, Right, Left
func NewMovement(maxVel [4]float64, accel [4]float64, decel [4]float64, enRate [4]float64, enBase [4]float64, turnBase float64, turnRate float64) *Movement {
	m := MoveProp{maxVel, accel, decel, enRate, enBase, turnRate, turnBase}
	return &Movement{Prop: &m}
}

// Turn rate used for setting arc for current direction
func (m *Movement) TurnRate() float64 {
	return m.Prop.TurnRateAt(fl.Norm(m.curVelocity[:], 2))
}

func (m *Movement) Command() ([2]float64, float64) {
	t := m.cmdVelocity[:]
	speed := fl.Norm(t, 2)
	fl.Scale(1/speed, t)

	return [2]float64{t[0], t[1]}, speed
}

// dir[1] > 0 ? Front : Back; dir[0] > 0 ? Right : Left
func (m *Movement) SetCommand(dir [2]float64, speed float64) {
	// Set to unit vector if not already
	fl.Scale(1/fl.Norm(dir[:], 2), dir[:])

	// Check max horizantal velocity
	hsp := dir[0]
	if hsp > 0 {
		hsp *= m.Prop.MaxVelocity[2]
	} else {
		hsp *= m.Prop.MaxVelocity[3]
	}

	// Check max vertical max velocity
	vsp := dir[1]
	if vsp > 0 {
		vsp *= m.Prop.MaxVelocity[0]
	} else {
		vsp *= m.Prop.MaxVelocity[1]
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

func (m *Movement) Update(del float64) float64 {
	return 0
}
