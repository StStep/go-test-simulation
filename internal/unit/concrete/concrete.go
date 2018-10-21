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
	db            ledger.LedgerRO
	prop          pr.Prop
	id            id.Uid
	formation     *form.Formation
	entityCommand chan int
	guide         id.Eid
	members       []id.Eid
}

func NewConstructor(db ledger.LedgerRO, conf conf.Configuration,
	idgen id.UidGen, entConstr ent.EntityConstructor) unit.UnitConstructor {
	return &constructor{db: db, conf: conf, idgen: idgen, entConstr: entConstr}
}

func (c *constructor) New(name string, pos [2]float64) (unit.Unit, []ent.Entity) {
	prop := c.conf.Unit(name)
	u := concrete{
		db:            c.db,
		prop:          prop,
		id:            c.idgen.Id(),
		formation:     form.NewFormation(c.conf.Formation(prop.Formations()[0]), prop.Size()),
		entityCommand: make(chan int),
		guide:         0,
		members:       make([]id.Eid, prop.Size()),
	}

	// TODO pull info from Formation, currently making block with leader at top left
	k := 0
	entRet := make([]ent.Entity, prop.Size())
	for name, count := range u.Prop().Members() {
		for i := 0; i < count; i++ {
			startPos := [2]float64{pos[0] + float64(k)*2.0, pos[1]}
			formOffset := [2]float64{float64(k % 5), float64(k / 5)}
			entRet[k] = c.entConstr.New(name, u.Id(), u.entityCommand, startPos, formOffset)
			u.members[k] = entRet[k].Id()
			if k == 0 {
				u.guide = u.members[k]
			}
			k++
		}
	}

	return &u, entRet
}

func (u *concrete) Id() id.Uid {
	return u.id
}

func (u *concrete) Prop() pr.Prop {
	return u.prop
}

func (u *concrete) Size() int {
	return len(u.members)
}

func (u *concrete) Guide() [2]float64 {
	return u.db.Entity(u.guide).Position()
}

func (e *concrete) LogicStep(del float64) {
	for _, v := range e.members {
		e.db.Entity(v).LogicStep(del)
	}
}
