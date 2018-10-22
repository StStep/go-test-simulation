package prop

type Prop interface {
	MaxVelocity() [4]float64
	Acceleration() [4]float64
	Deceleration() [4]float64
	EnergyUsageRate() [4]float64
	BaseEnergyUsage() [4]float64
	TurnRadiusRate() float64
	BaseTurnRadius() float64
	FootprintRadius() float64
	TurnRateAt(speed float64) float64
}

type prop struct {
	maxVelocity                      [4]float64
	acceleration, deceleration       [4]float64
	energyUsageRate, baseEnergyUsage [4]float64
	turnRadiusRate, baseTurnRadius   float64
	footprintRadius                  float64
}

func New(maxVel [4]float64, accel [4]float64, decel [4]float64, enRate [4]float64, enBase [4]float64, turnRate float64, turnBase float64, radius float64) Prop {
	return &prop{maxVel, accel, decel, enRate, enBase, turnRate, turnBase, radius}
}

func (p *prop) MaxVelocity() [4]float64 {
	return p.maxVelocity
}

func (p *prop) Acceleration() [4]float64 {
	return p.acceleration
}

func (p *prop) Deceleration() [4]float64 {
	return p.deceleration
}

func (p *prop) EnergyUsageRate() [4]float64 {
	return p.energyUsageRate
}

func (p *prop) BaseEnergyUsage() [4]float64 {
	return p.baseEnergyUsage
}

func (p *prop) TurnRadiusRate() float64 {
	return p.turnRadiusRate
}

func (p *prop) BaseTurnRadius() float64 {
	return p.baseTurnRadius
}

func (p *prop) FootprintRadius() float64 {
	return p.footprintRadius
}

func (p *prop) TurnRateAt(speed float64) float64 {
	ret := p.BaseTurnRadius() + p.TurnRadiusRate()*speed
	if ret < 0 {
		return 0
	} else {
		return ret
	}
}
