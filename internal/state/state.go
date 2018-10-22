package state

import (
	conf "github.com/StStep/go-test-simulation/internal/configuration"
	ent "github.com/StStep/go-test-simulation/internal/entity"
	"github.com/StStep/go-test-simulation/internal/id"
	"github.com/StStep/go-test-simulation/internal/physics"
	un "github.com/StStep/go-test-simulation/internal/unit"
)

type State struct {
	Configuration *conf.Configuration
	UidGen        *id.UidGen
	EidGen        *id.EidGen
	Ledger        *Ledger
	Physics       *physics.Physics
}

type Ledger struct {
	UnitData   map[id.Uid]*un.Unit
	EntityData map[id.Eid]*ent.Entity
}
