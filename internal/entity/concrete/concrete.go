package concrete

import (
	"github.com/StStep/go-test-simulation/internal/entity"
	pr "github.com/StStep/go-test-simulation/internal/entity/prop"
	"github.com/StStep/go-test-simulation/internal/id"
	mv "github.com/StStep/go-test-simulation/internal/movement"
	sp "github.com/StStep/go-test-simulation/internal/space"
)

type concrete struct {
	prop         pr.Prop
	id           id.Eid
	movement     *mv.Movement
	spaceViewer  sp.SpaceViewer  // Use to See others in same space
	spaceUpdater sp.SpaceUpdater // Use to move self around
	command      chan int
	formOffset   [2]float64
}

func (e *concrete) Prop() pr.Prop {
	return e.prop
}

func (e *concrete) Id() id.Eid {
	return e.id
}

func NewEntity(prop pr.Prop, cmd chan int) entity.Entity {
	return &concrete{
		prop:         prop,
		movement:     mv.NewMovement(prop.Movement()),
		spaceViewer:  nil,
		spaceUpdater: nil,
		command:      cmd,
		formOffset:   [2]float64{0, 0},
	}
}

// TODO Use SpaceUpdater and SpaceViewer in update step to follow CmdReform
