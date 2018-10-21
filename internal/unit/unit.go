package unit

import (
	ent "github.com/StStep/go-test-simulation/internal/entity"
	"github.com/StStep/go-test-simulation/internal/id"
	pr "github.com/StStep/go-test-simulation/internal/unit/prop"
)

type UnitConstructor interface {
	New(name string, pos [2]float64) (Unit, []ent.Entity)
}

type Unit interface {
	Id() id.Uid
	Prop() pr.Prop
	Size() int
	Guide() [2]float64
	LogicStep(del float64)
}
