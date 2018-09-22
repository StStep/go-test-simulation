package movement

import (
	fl "gonum.org/v1/gonum/floats"
	"math"
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

type Movement struct {
	Properties   *MoveProp  // Movement properties to use with math
	CurVelocity  [2]float64 // Represents current velocity vector
	CmdDirection [2]float64 // Unit vector for desired movement direction
	CmdSpeed     float64    // Scaler representing desired movement speed
}

// Forward Backward, Right, Left
func NewMovement(maxVel [4]float64) *Movement {
	m := MoveProp{
		Forward:  MoveDirProp{MaxVelocity: maxVel[0]},
		Backward: MoveDirProp{MaxVelocity: maxVel[1]},
		Right:    MoveDirProp{MaxVelocity: maxVel[2]},
		Left:     MoveDirProp{MaxVelocity: maxVel[3]},
	}
	return &Movement{Properties: &m}
}

// dir[1] > 0 ? Front : Back; dir[0] > 0 ? Right : Left
func (m *Movement) SetCommand(dir [2]float64, speed float64) {
	// Set to unit vector if not already
	fl.Scale(1/fl.Norm(dir[:], 1), dir[:])
	copy(m.CmdDirection[:], dir[:])

	// Check horizantal velocity
	hsp := 0.0
	if m.CmdDirection[0] > 0 {
		hsp = math.Min(speed, m.Properties.Right.MaxVelocity) * m.CmdDirection[0]
	} else {
		hsp = math.Min(speed, m.Properties.Left.MaxVelocity) * m.CmdDirection[0] * -1
	}

	// Check vertical max velocity
	vsp := 0.0
	if m.CmdDirection[1] > 0 {
		vsp = math.Min(speed, m.Properties.Forward.MaxVelocity) * m.CmdDirection[1]
	} else {
		vsp = math.Min(speed, m.Properties.Backward.MaxVelocity) * m.CmdDirection[1] * -1
	}

	// Combine
	m.CmdSpeed = vsp + hsp
}

func (m *Movement) Update(del float64) float64 {
	return 0
}
