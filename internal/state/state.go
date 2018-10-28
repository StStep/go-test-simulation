package state

import (
	"github.com/StStep/go-test-simulation/internal/physics"
	conf "github.com/StStep/go-test-simulation/internal/state/configuration"
	unit "github.com/StStep/go-test-simulation/internal/unit"
)

type State struct {
	Configuration *conf.Configuration
	Ledger        *Ledger
	Physics       *physics.Physics
	uIdGen        *idGen
	eidGen        *idGen
}

type Ledger struct {
	UnitData   map[uint64]*unit.Unit
	EntityData map[uint64]*unit.Entity
}

type idGen struct {
	lastid uint64
}

func (gen *idGen) Reset() {
	gen.lastid = 0
}

func (gen *idGen) Id() uint64 {
	gen.lastid = gen.lastid + 1
	return gen.lastid
}

func (state *State) newEntity(name string, uid uint64, cmd chan int, pos [2]float64, offset [2]float64) *unit.Entity {
	//	prop := state.Configuration.Entities[name]
	//	e := ent.Entity{
	//		Id:         state.EidGen.Id(),
	//		Prop:       prop,
	//		Unit:       uid,
	//		Command:    cmd,
	//		FormOffset: offset,
	//	}
	//	state.Physics.RegisterEntity(e.Id, prop.Physics, pos)
	//	return &e
	return nil
}

func (state *State) NewUnit(name string, pos [2]float64) {
	//	prop := state.Configuration.Units[name]
	//	u := un.Unit{
	//		Id:            state.Uidgen.Id(),
	//		Prop:          prop,
	//		Formation:     form.NewFormation(c.conf.Formation(prop.Formations()[0]), prop.Size()),
	//		EntityCommand: make(chan int),
	//		Guide:         0,
	//		Members:       make([]id.Eid, prop.Size()),
	//	}
	//
	//	// TODO pull info from Formation, currently making block with leader at top left
	//	k := 0
	//	entRet := make([]ent.Entity, prop.Size())
	//	for name, count := range u.Prop().Members() {
	//		for i := 0; i < count; i++ {
	//			startPos := [2]float64{pos[0] + float64(k)*2.0, pos[1]}
	//			formOffset := [2]float64{float64(k % 5), float64(k / 5)}
	//			entRet[k] = c.entConstr.New(name, u.Id(), u.entityCommand, startPos, formOffset)
	//			u.members[k] = entRet[k].Id()
	//			if k == 0 {
	//				u.guide = u.members[k]
	//			}
	//			k++
	//		}
	//	}
	//
	//	return &u, entRet
}
