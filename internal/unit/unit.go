package unit

import (
	cf "github.com/StStep/go-test-simulation/internal/configuration"
	ent "github.com/StStep/go-test-simulation/internal/entity"
	form "github.com/StStep/go-test-simulation/internal/formation"
	sp "github.com/StStep/go-test-simulation/internal/space"
	pr "github.com/StStep/go-test-simulation/internal/unitprop"
)

const (
	CmdNil = 0
)

type Unit struct {
	Prop          pr.UnitProp
	Formation     *form.Formation
	EntityCommand chan int
	Leader        *ent.Entity
	Members       []*ent.Entity
}

func NewUnit(prop pr.UnitProp, pos [2]float64, space *sp.Space, conf cf.Configuration) *Unit {
	u := Unit{Prop: prop}
	u.Formation = form.NewFormation(conf.Formation(u.Prop.Formations()[0]), u.Prop.Size())
	u.EntityCommand = make(chan int)
	if u.Prop.Leader() != "" {
		u.Leader = ent.NewEntity(conf.Entity(u.Prop.Leader()), u.EntityCommand)
		u.Leader.SpaceViewer, u.Leader.SpaceUpdater = space.Register(pos, u.Leader.Prop.Radius())
		u.Leader.FormOffset = [2]float64{0, 0}
	} else {
		u.Leader = nil
	}
	u.Members = make([]*ent.Entity, u.Prop.Size())
	for i := 0; i < u.Prop.Size(); i++ {
		u.Members[i] = ent.NewEntity(conf.Entity(u.Prop.Members()[i]), u.EntityCommand)
		startPos := [2]float64{pos[0] + u.Members[i].Prop.Radius()*4, pos[1]}
		u.Members[i].SpaceViewer, u.Members[i].SpaceUpdater =
			space.Register(startPos, u.Members[i].Prop.Radius())
		// TODO pull info from Formation, currently making block with leader at top left
		u.Members[i].FormOffset = [2]float64{float64((i + 1) % 5), float64((i + 1) / 5)}
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

// TODO Only give reform command to members
func (u *Unit) GiveCommand(cmd int) {
	u.EntityCommand <- ent.CmdReform
}
