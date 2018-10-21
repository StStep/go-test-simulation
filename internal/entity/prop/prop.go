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

type prop struct {
	name   string
	phy    phy.Prop
	radius float64
}

func NewProp(name string, phy phy.Prop, radius float64) Prop {
	return &prop{name, phy, radius}
}

func (p *prop) Name() string {
	return p.name
}

func (p *prop) Physics() phy.Prop {
	return p.phy
}

func (p *prop) Radius() float64 {
	return p.radius
}
