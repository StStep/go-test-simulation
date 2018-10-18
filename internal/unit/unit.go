package unit

import (
	"github.com/StStep/go-test-simulation/internal/id"
)

type UnitConstructor interface {
	New(name string, pos [2]float64) Unit
}

type Unit interface {
	Id() id.Uid
}
