package prop

type Prop struct {
	MaxVelocity     [4]float64 `json:"maxVelocity"`
	Acceleration    [4]float64 `json:"acceleration"`
	Deceleration    [4]float64 `json:"deceleration"`
	EnergyUsageRate [4]float64 `json:"energyUsageRate"`
	BaseEnergyUsage [4]float64 `json:"baseEnergyUsage"`
	TurnRadiusRate  float64    `json:"turnRadiusRate"`
	BaseTurnRadius  float64    `json:"baseTurnRadius"`
	FootprintRadius float64    `json:"footprintRadius"`
}

func (p *Prop) TurnRateAt(speed float64) float64 {
	ret := p.BaseTurnRadius + p.TurnRadiusRate*speed
	if ret < 0 {
		return 0
	} else {
		return ret
	}
}
