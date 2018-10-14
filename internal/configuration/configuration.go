package configuration

import (
	entpr "github.com/StStep/go-test-simulation/internal/entityprop"
	forpr "github.com/StStep/go-test-simulation/internal/formationprop"
)

type Configuration interface {
	Entity(id string) entpr.EntityProp
	Formation(id string) forpr.FormationProp
}

func NewConfiguration() Configuration {
	c := conf{}
	c.entities = make(map[string]entpr.EntityProp)
	c.formations = make(map[string]forpr.FormationProp)
	return &c
}

type conf struct {
	entities   map[string]entpr.EntityProp
	formations map[string]forpr.FormationProp
}

func (c *conf) Entity(id string) entpr.EntityProp {
	return c.entities[id]
}

func (c *conf) Formation(id string) forpr.FormationProp {
	return c.formations[id]
}
