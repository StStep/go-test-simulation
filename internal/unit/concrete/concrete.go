package concrete

import (
	conf "github.com/StStep/go-test-simulation/internal/configuration"
	ent "github.com/StStep/go-test-simulation/internal/entity"
	form "github.com/StStep/go-test-simulation/internal/formation"
	"github.com/StStep/go-test-simulation/internal/id"
	"github.com/StStep/go-test-simulation/internal/ledger"
	"github.com/StStep/go-test-simulation/internal/unit"
	pr "github.com/StStep/go-test-simulation/internal/unit/prop"
)

const (
	CmdNil = 0
)

type constructor struct {
	db        ledger.LedgerRO
	conf      conf.Configuration
	idgen     id.UidGen
	entConstr ent.EntityConstructor
}

type concrete struct {
	prop          pr.Prop
	id            id.Uid
	formation     *form.Formation
	entityCommand chan int
	leader        ent.Entity
	members       []ent.Entity
}

func NewConstructor(db ledger.LedgerRO, conf conf.Configuration,
	idgen id.UidGen, entConstr ent.EntityConstructor) unit.UnitConstructor {
	return &constructor{db: db, conf: conf, idgen: idgen, entConstr: entConstr}
}

func (c *constructor) New(name string, pos [2]float64) unit.Unit {
	prop := c.conf.Unit(name)
	u := concrete{
		prop:          prop,
		id:            c.idgen.Id(),
		formation:     form.NewFormation(c.conf.Formation(prop.Formations()[0]), prop.Size()),
		entityCommand: make(chan int),
		leader:        nil,
		members:       make([]ent.Entity, prop.Size()),
	}

	// Leader
	if u.prop.Leader() != "" {
		u.leader = c.entConstr.New(u.prop.Leader(), u.entityCommand, pos, [2]float64{0, 0})
	}

	// Member
	for i := 0; i < u.prop.Size(); i++ {
		// TODO pull info from Formation, currently making block with leader at top left
		startPos := [2]float64{pos[0] + float64(i)*2.0, pos[1]}
		formOffset := [2]float64{float64((i + 1) % 5), float64((i + 1) / 5)}
		u.members[i] = c.entConstr.New(u.prop.Members()[i], u.entityCommand, startPos, formOffset)
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
