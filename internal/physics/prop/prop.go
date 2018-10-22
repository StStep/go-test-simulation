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

type Pprop struct {
	PmaxVelocity     [4]float64 `json:"maxVelocity"`
	Pacceleration    [4]float64 `json:"acceleration"`
	Pdeceleration    [4]float64 `json:"deceleration"`
	PenergyUsageRate [4]float64 `json:"energyUsageRate"`
	PbaseEnergyUsage [4]float64 `json:"baseEnergyUsage"`
	PturnRadiusRate  float64    `json:"turnRadiusRate"`
	PbaseTurnRadius  float64    `json:"baseTurnRadius"`
	PfootprintRadius float64    `json:"footprintRadius"`
}

func New(maxVel [4]float64, accel [4]float64, decel [4]float64, enRate [4]float64, enBase [4]float64, turnRate float64, turnBase float64, radius float64) *Pprop {
	return &Pprop{maxVel, accel, decel, enRate, enBase, turnRate, turnBase, radius}
}

func (p *Pprop) MaxVelocity() [4]float64 {
	return p.PmaxVelocity
}

func (p *Pprop) Acceleration() [4]float64 {
	return p.Pacceleration
}

func (p *Pprop) Deceleration() [4]float64 {
	return p.Pdeceleration
}

func (p *Pprop) EnergyUsageRate() [4]float64 {
	return p.PenergyUsageRate
}

func (p *Pprop) BaseEnergyUsage() [4]float64 {
	return p.PbaseEnergyUsage
}

func (p *Pprop) TurnRadiusRate() float64 {
	return p.PturnRadiusRate
}

func (p *Pprop) BaseTurnRadius() float64 {
	return p.PbaseTurnRadius
}

func (p *Pprop) FootprintRadius() float64 {
	return p.PfootprintRadius
}

func (p *Pprop) TurnRateAt(speed float64) float64 {
	ret := p.BaseTurnRadius() + p.TurnRadiusRate()*speed
	if ret < 0 {
		return 0
	} else {
		return ret
	}
}
