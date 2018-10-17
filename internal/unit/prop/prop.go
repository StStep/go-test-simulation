package unitprop

import ()

type Prop interface {
	Name() string
	Leader() string
	Members() []string
	Formations() []string
	Size() int
}

type prop struct {
	name                string
	leader              string
	members, formations []string
}

func NewProp(name string, leader string, members []string, formations []string) Prop {
	return &prop{name, leader, members, formations}
}

func (p *prop) Name() string {
	return p.name
}

func (p *prop) Leader() string {
	return p.leader
}

func (p *prop) Members() []string {
	return p.members
}

func (p *prop) Formations() []string {
	return p.formations
}

func (p *prop) Size() int {
	base := len(p.Members())
	if p.Leader() != "" {
		return base + 1
	} else {
		return base
	}
}
