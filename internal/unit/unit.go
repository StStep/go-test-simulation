package unit

import (
	cf "github.com/StStep/go-test-simulation/internal/configuration"
	ent "github.com/StStep/go-test-simulation/internal/entity"
	form "github.com/StStep/go-test-simulation/internal/formation"
	sp "github.com/StStep/go-test-simulation/internal/space"
	pr "github.com/StStep/go-test-simulation/internal/unitprop"
)

type Unit struct {
	Prop          pr.UnitProp
	Formation     *form.Formation
	UnitCommand   chan int
	EntityCommand chan int
	Leader        *ent.Entity
	Members       []*ent.Entity
}

func NewUnit(prop pr.UnitProp, cmd chan int, pos [2]float64, space *sp.Space, conf cf.Configuration) *Unit {
	u := Unit{Prop: prop}
	u.Formation = form.NewFormation(conf.Formation(u.Prop.Formations()[0]), u.Prop.Size())
	u.UnitCommand = cmd
	u.EntityCommand = make(chan int)
	if u.Prop.Leader() != "" {
		u.Leader = ent.NewEntity(conf.Entity(u.Prop.Leader()), u.EntityCommand)
		u.Leader.SpaceId = space.RegisterEntity(pos, u.Leader.Prop.Radius())
	} else {
		u.Leader = nil
	}
	u.Members = make([]*ent.Entity, u.Prop.Size())
	for i := 0; i < u.Prop.Size(); i++ {
		u.Members[i] = ent.NewEntity(conf.Entity(u.Prop.Members()[i]), u.EntityCommand)
		startPos := [2]float64{pos[0] + u.Members[i].Prop.Radius()*4, pos[1]}
		u.Members[i].SpaceId = space.RegisterEntity(startPos, u.Members[i].Prop.Radius())
	}
	return &u
}

func (u *Unit) Size() int {
	base := len(u.Members)
	if u.Leader != nil {
		return base + 1
	} else {
		return base
	}
}

func (u *Unit) UpdateMove(done chan bool) {
	done <- false
}
