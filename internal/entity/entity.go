package entity

import (
	pr "github.com/StStep/go-test-simulation/internal/entityprop"
	mv "github.com/StStep/go-test-simulation/internal/movement"
	sp "github.com/StStep/go-test-simulation/internal/space"
)

const (
	CmdNil    = 0
	CmdReform = 1
)

type Entity struct {
	Prop         pr.EntityProp
	Movement     *mv.Movement
	SpaceViewer  sp.SpaceViewer  // Use to See others in same space
	SpaceUpdater sp.SpaceUpdater // Use to move self around
	Command      chan int
	FormOffset   [2]float64
}

func NewEntity(prop pr.EntityProp, cmd chan int) *Entity {
	return &Entity{
		Prop:         prop,
		Movement:     mv.NewMovement(prop.Movement()),
		SpaceViewer:  nil,
		SpaceUpdater: nil,
		Command:      cmd,
		FormOffset:   [2]float64{0, 0},
	}
}
