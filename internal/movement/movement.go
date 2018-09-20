package movement

import (
	"github.com/StStep/go-test-simulation/internal/vecmath"
)

type Direction int

const (
	DirNone     Direction = 0
	DirForward  Direction = 1
	DirBackward Direction = 2
	DirLeft     Direction = 3
	DirRight    Direction = 4
)

type MoveDirProp struct {
	MaxVelocity                      float64
	Acceleration, Deceleration       float64
	EnergyUsageRate, BaseEnergyUsage float64
}

type MoveProp struct {
	Forward, Backward, Left, Right MoveDirProp
	TurnRadiusRate, BaseTurnRadius float64
}

type Movement struct {
	Properties   *MoveProp
	CurVelocity  vecmath.Vector
	CmdDirection Direction
	CmdSpeed     float64
}

func (l *Movement) Update(del float64) float64 {
	return 0
}
