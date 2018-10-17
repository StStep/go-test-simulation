package entity

import (
	"github.com/StStep/go-test-simulation/internal/entity/prop"
	"github.com/StStep/go-test-simulation/internal/id"
)

const (
	CmdNil    = 0
	CmdReform = 1
)

type Entity interface {
	Prop() prop.Prop
	Id() id.Eid
}
