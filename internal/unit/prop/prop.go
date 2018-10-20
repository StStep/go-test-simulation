package unitprop

import ()

type Prop interface {
	Name() string
	Members() map[string]int
	Formations() []string
	Size() int
}

type prop struct {
	name       string
	leader     string
	members    map[string]int
	formations []string
}

func NewProp(name string, leader string, members map[string]int, formations []string) Prop {
	return &prop{name, leader, members, formations}
}

func (p *prop) Name() string {
	return p.name
}

func (p *prop) Members() map[string]int {
	return p.members
}

func (p *prop) Formations() []string {
	return p.formations
}

func (p *prop) Size() int {
	ret := 0
	for _, v := range p.Members() {
		ret = ret + v
	}
	return ret
}
