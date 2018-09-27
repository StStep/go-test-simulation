package movement

import (
	fl "gonum.org/v1/gonum/floats"
)

type MoveDirProp struct {
	MaxVelocity                      float64
	Acceleration, Deceleration       float64
	EnergyUsageRate, BaseEnergyUsage float64
}

type MoveProp struct {
	Forward, Backward, Right, Left MoveDirProp
	TurnRadiusRate, BaseTurnRadius float64
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
	Properties  *MoveProp  // Movement properties to use with math
	curVelocity [2]float64 // Represents current velocity vector
	cmdVelocity [2]float64 // Represents commanded velocity vector
}

// Forward Backward, Right, Left
func NewMovement(maxVel [4]float64, accel [4]float64, decel [4]float64, enRate [4]float64, enBase [4]float64, turnBase float64, turnRate float64) *Movement {
	m := MoveProp{
		Forward:        MoveDirProp{MaxVelocity: maxVel[0], Acceleration: accel[0], Deceleration: decel[0], EnergyUsageRate: enRate[0], BaseEnergyUsage: enBase[0]},
		Backward:       MoveDirProp{MaxVelocity: maxVel[1], Acceleration: accel[1], Deceleration: decel[1], EnergyUsageRate: enRate[1], BaseEnergyUsage: enBase[1]},
		Right:          MoveDirProp{MaxVelocity: maxVel[2], Acceleration: accel[2], Deceleration: decel[2], EnergyUsageRate: enRate[2], BaseEnergyUsage: enBase[2]},
		Left:           MoveDirProp{MaxVelocity: maxVel[3], Acceleration: accel[3], Deceleration: decel[3], EnergyUsageRate: enRate[3], BaseEnergyUsage: enBase[3]},
		TurnRadiusRate: turnRate,
		BaseTurnRadius: turnBase,
	}
	return &Movement{Properties: &m}
}

// Turn rate used for setting arc for current direction
func (m *Movement) TurnRate() float64 {
	return m.Properties.TurnRateAt(fl.Norm(m.curVelocity[:], 2))
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
		hsp *= m.Properties.Right.MaxVelocity
	} else {
		hsp *= m.Properties.Left.MaxVelocity
	}

	// Check max vertical max velocity
	vsp := dir[1]
	if vsp > 0 {
		vsp *= m.Properties.Forward.MaxVelocity
	} else {
		vsp *= m.Properties.Backward.MaxVelocity
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
