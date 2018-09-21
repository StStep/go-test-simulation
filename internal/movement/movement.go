package movement

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
	Properties   *MoveProp  // Movement properties to use with math
	CurVelocity  [2]float64 // Represents current velocity vector
	CmdDirection [2]float64 // Unit vector for desired movement direction
	CmdSpeed     float64    // Scaler representing desired movement speed
}

func (m *Movement) SetCommand(dir [2]float64, speed float64) {

}

func (m *Movement) Update(del float64) float64 {
	return 0
}
