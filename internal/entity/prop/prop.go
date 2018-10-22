package prop

import (
	phy "github.com/StStep/go-test-simulation/internal/physics/prop"
)

// TODO Could decorate props with modifiers? Or decorate with a sturct allowing modifiers?
type Prop interface {
	Name() string
	Physics() phy.Prop
	Radius() float64
}

type Pprop struct {
	Pname string     `json:"name"`
	Pphy  *phy.Pprop `json:"physics"`
}

func New(name string, phy *phy.Pprop) *Pprop {
	return &Pprop{name, phy}
}

func (p *Pprop) Name() string {
	return p.Pname
}

func (p *Pprop) Physics() phy.Prop {
	return p.Pphy
}

func (p *Pprop) Radius() float64 {
	return p.Pphy.FootprintRadius()
}
