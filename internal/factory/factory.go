package factory

import (
	conf "github.com/StStep/go-test-simulation/internal/configuration"
	entc "github.com/StStep/go-test-simulation/internal/entity/concrete"
	"github.com/StStep/go-test-simulation/internal/id"
	"github.com/StStep/go-test-simulation/internal/ledger"
	"github.com/StStep/go-test-simulation/internal/physics"
	unit "github.com/StStep/go-test-simulation/internal/unit"
	unitc "github.com/StStep/go-test-simulation/internal/unit/concrete"
)

func New(db ledger.LedgerRO, conf conf.Configuration, phy physics.Physics, uidgen id.UidGen, eidgen id.EidGen) unit.UnitConstructor {
	econstr := entc.NewConstructor(db, conf, eidgen, phy)
	uconstr := unitc.NewConstructor(db, conf, uidgen, econstr)
	return uconstr
}
