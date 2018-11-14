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
	eIdGen        *idGen
}

func New(confFile string) *State {
	state := State{conf.FromFile(confFile), newLedger(), physics.New(), newIdGen(), newIdGen()}
	return &state
}

func (state *State) newEntity(name string, uid uint64, cmd chan int, pos [2]float64, offset [2]float64) uint64 {
	prop := state.Configuration.Entities[name]
	e := unit.Entity{
		Id:         state.eIdGen.Id(),
		Prop:       prop,
		UnitId:     uid,
		Command:    cmd,
		FormOffset: offset,
	}
	state.Ledger.EntityData[e.Id] = &e
	state.Physics.RegisterEntity(e.Id, prop.Physics, pos)
	return e.Id
}

func (state *State) NewUnit(name string, pos [2]float64) uint64 {
	prop := state.Configuration.Units[name]
	u := unit.Unit{
		Id:            state.uIdGen.Id(),
		Prop:          prop,
		GuideId:       0,
		MemberIds:     make([]uint64, prop.Size()),
		EntityCommand: make(chan int),
	}

	// TODO pull info from Formation, currently making block with leader at top left
	k := 0
	for name, count := range u.Prop.Members {
		for i := 0; i < count; i++ {
			startPos := [2]float64{pos[0] + float64(k)*2.0, pos[1]}
			formOffset := [2]float64{float64(k % 5), float64(k / 5)}
			u.MemberIds[k] = state.newEntity(name, u.Id, u.EntityCommand, startPos, formOffset)
			if k == 0 {
				u.GuideId = state.Ledger.EntityData[u.MemberIds[k]].Id
			}
			k++
		}
	}
	return u.Id
}

type Ledger struct {
	UnitData   map[uint64]*unit.Unit
	EntityData map[uint64]*unit.Entity
}

func newLedger() *Ledger {
	return &Ledger{make(map[uint64]*unit.Unit), make(map[uint64]*unit.Entity)}
}

type idGen struct {
	lastid uint64
}

func newIdGen() *idGen {
	gen := idGen{}
	gen.Reset()
	return &gen
}

func (gen *idGen) Reset() {
	gen.lastid = 0
}

func (gen *idGen) Id() uint64 {
	gen.lastid = gen.lastid + 1
	return gen.lastid
}
