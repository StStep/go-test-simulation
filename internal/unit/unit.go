package unit

import (
	ent "github.com/StStep/go-test-simulation/internal/entity"
	"github.com/StStep/go-test-simulation/internal/id"
)

type UnitConstructor interface {
	New(name string, pos [2]float64) (Unit, []ent.Entity)
}

type Unit struct {
	Id      id.Uid
	Prop    Prop
	Guide   id.Eid
	Members []id.Eid
}

func (u *Unit) Size() int {
	return len(u.Members)
}

type Prop struct {
	Name       string         `json:"name"`
	Members    map[string]int `json:"members"`
	Formations []string       `json:"formations"`
}

func (p *Prop) Size() int {
	ret := 0
	for _, v := range p.Members {
		ret = ret + v
	}
	return ret
}
