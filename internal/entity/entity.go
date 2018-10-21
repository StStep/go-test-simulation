package entity

import (
	"github.com/StStep/go-test-simulation/internal/entity/prop"
	"github.com/StStep/go-test-simulation/internal/id"
)

const (
	CmdNil    = 0
	CmdReform = 1
)

type EntityConstructor interface {
	New(name string, uid id.Uid, cmd chan int, pos [2]float64, offset [2]float64) Entity
}

type Entity interface {
	Id() id.Eid
	Prop() prop.Prop
	Position() [2]float64
	Velocity() [2]float64
}
