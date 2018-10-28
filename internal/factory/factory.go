package factory

import (
	_ "github.com/StStep/go-test-simulation/internal/configuration"
	_ "github.com/StStep/go-test-simulation/internal/physics"
	"github.com/StStep/go-test-simulation/internal/state"
	"github.com/StStep/go-test-simulation/internal/unit"
)

func NewEntity(state *state.State, name string, uid uint64, cmd chan int, pos [2]float64, offset [2]float64) *unit.Entity {
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

func NewUnit(state *state.State, name string, pos [2]float64) (unit.Unit, []unit.Entity) {
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
	return unit.Unit{}, nil
}
