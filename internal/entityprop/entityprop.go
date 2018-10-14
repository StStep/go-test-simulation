package entityprop

import (
	mpr "github.com/StStep/go-test-simulation/internal/movementprop"
)

// TODO Could decorate props with modifiers? Or decorate with a sturct allowing modifiers?
type EntityProp interface {
	Name() string
	Movement() mpr.MovementProp
	Radius() float64
}

type prop struct {
	name     string
	movement mpr.MovementProp
	radius   float64
}

func NewEntityProp(name string, movement mpr.MovementProp, radius float64) EntityProp {
	return &prop{name, movement, radius}
}

func (p *prop) Name() string {
	return p.name
}

func (p *prop) Movement() mpr.MovementProp {
	return p.movement
}

func (p *prop) Radius() float64 {
	return p.radius
}
