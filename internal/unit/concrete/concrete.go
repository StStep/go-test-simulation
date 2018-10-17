package concrete

import (
	cf "github.com/StStep/go-test-simulation/internal/configuration"
	ent "github.com/StStep/go-test-simulation/internal/entity"
	form "github.com/StStep/go-test-simulation/internal/formation"
	"github.com/StStep/go-test-simulation/internal/id"
	sp "github.com/StStep/go-test-simulation/internal/space"
	"github.com/StStep/go-test-simulation/internal/unit"
	pr "github.com/StStep/go-test-simulation/internal/unit/prop"
)

const (
	CmdNil = 0
)

type concrete struct {
	prop          pr.Prop
	id            id.Uid
	formation     *form.Formation
	entityCommand chan int
	leader        ent.Entity
	members       []ent.Entity
}

func NewUnit(prop pr.Prop, pos [2]float64, space *sp.Space, conf cf.Configuration) unit.Unit {
	u := concrete{prop: prop}
	u.formation = form.NewFormation(conf.Formation(u.prop.Formations()[0]), u.prop.Size())
	u.entityCommand = make(chan int)
	if u.prop.Leader() != "" {
		//u.leader = ent.NewEntity(conf.Entity(u.prop.Leader()), u.EntityCommand)
		//u.leader.SpaceViewer, u.Leader.SpaceUpdater = space.Register(pos, u.Leader.Prop.Radius())
		//u.leader.FormOffset = [2]float64{0, 0}
	} else {
		u.leader = nil
	}
	u.members = make([]ent.Entity, u.prop.Size())
	for i := 0; i < u.prop.Size(); i++ {
		//u.Members[i] = ent.NewEntity(conf.Entity(u.Prop.Members()[i]), u.EntityCommand)
		//startPos := [2]float64{pos[0] + u.Members[i].Prop.Radius()*4, pos[1]}
		//u.Members[i].SpaceViewer, u.Members[i].SpaceUpdater =
		//	space.Register(startPos, u.Members[i].Prop.Radius())
		//// TODO pull info from Formation, currently making block with leader at top left
		//u.Members[i].FormOffset = [2]float64{float64((i + 1) % 5), float64((i + 1) / 5)}
	}
	return &u
}

func (u *concrete) Id() id.Uid {
	return u.id
}

func (u *concrete) Size() int {
	base := len(u.members)
	if u.leader != nil {
		return base + 1
	} else {
		return base
	}
}

// TODO Only give reform command to members
func (u *concrete) GiveCommand(cmd int) {
	u.entityCommand <- ent.CmdReform
}
