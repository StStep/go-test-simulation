package ledger

import (
	ent "github.com/StStep/go-test-simulation/internal/entity"
	"github.com/StStep/go-test-simulation/internal/id"
	un "github.com/StStep/go-test-simulation/internal/unit"
)

type LedgerRO interface {
	Unit(id id.Uid) un.Unit
	Entity(id id.Eid) ent.Entity
}

type Ledger interface {
	Unit(id id.Uid) un.Unit
	Entity(id id.Eid) ent.Entity
	AddUnit(unit un.Unit)
	AddEntity(entity ent.Entity)
}

type ledgerdb struct {
	unitdb   map[id.Uid]un.Unit
	entitydb map[id.Eid]ent.Entity
}

func (l *ledgerdb) Unit(id id.Uid) un.Unit {
	return l.unitdb[id]
}

func (l *ledgerdb) Entity(id id.Eid) ent.Entity {
	return l.entitydb[id]
}

func (l *ledgerdb) AddUnit(unit un.Unit) {
	l.unitdb[unit.Id()] = unit
}

func (l *ledgerdb) AddEntity(entity ent.Entity) {
	l.entitydb[entity.Id()] = entity
}

func NewLedger() Ledger {
	return &ledgerdb{}
}
