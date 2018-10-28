package state

import (
	conf "github.com/StStep/go-test-simulation/internal/configuration"
	"github.com/StStep/go-test-simulation/internal/physics"
	unit "github.com/StStep/go-test-simulation/internal/unit"
)

type State struct {
	Configuration *conf.Configuration
	UIdGen        *IdGen
	EidGen        *IdGen
	Ledger        *Ledger
	Physics       *physics.Physics
}

type Ledger struct {
	UnitData   map[uint64]*unit.Unit
	EntityData map[uint64]*unit.Entity
}

func NewIdGen() *IdGen {
	return &IdGen{}
}

type IdGen struct {
	lastid uint64
}

func (gen *IdGen) Reset() {
	gen.lastid = 0
}

func (gen *IdGen) Id() uint64 {
	gen.lastid = gen.lastid + 1
	return gen.lastid
}
