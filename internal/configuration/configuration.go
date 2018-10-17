package configuration

import (
	ent "github.com/StStep/go-test-simulation/internal/entity/prop"
	un "github.com/StStep/go-test-simulation/internal/unit/prop"
	form "github.com/StStep/go-test-simulation/internal/formationprop"
)

type Configuration interface {
	Entity(name string) ent.Prop
	Unit(name string) un.Prop
	Formation(name string) form.FormationProp
}

func NewConfiguration() Configuration {
	c := conf{}
	c.entities = make(map[string]ent.Prop)
	c.formations = make(map[string]form.FormationProp)
	return &c
}

type conf struct {
	entities   map[string]ent.Prop
	units   map[string]un.Prop
	formations map[string]form.FormationProp
}

func (c *conf) Entity(name string) ent.Prop {
	return c.entities[name]
}

func (c *conf) Unit(name string) un.Prop {
	return c.units[name]
}

func (c *conf) Formation(name string) form.FormationProp {
	return c.formations[name]
}
