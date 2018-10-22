package prop

import ()

type Prop interface {
	Name() string
	Members() map[string]int
	Formations() []string
	Size() int
}

type Pprop struct {
	Pname       string         `json:"name"`
	Pmembers    map[string]int `json:"members"`
	Pformations []string       `json:"formations"`
}

func New(name string, members map[string]int, formations []string) *Pprop {
	return &Pprop{name, members, formations}
}

func (p *Pprop) Name() string {
	return p.Pname
}

func (p *Pprop) Members() map[string]int {
	return p.Pmembers
}

func (p *Pprop) Formations() []string {
	return p.Pformations
}

func (p *Pprop) Size() int {
	ret := 0
	for _, v := range p.Members() {
		ret = ret + v
	}
	return ret
}
