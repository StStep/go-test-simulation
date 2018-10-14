package entity

import (
	pr "github.com/StStep/go-test-simulation/internal/entityprop"
	mv "github.com/StStep/go-test-simulation/internal/movement"
)

type Entity struct {
	Prop     pr.EntityProp
	Movement *mv.Movement
	SpaceId  int
	Command  chan int
}

func NewEntity(prop pr.EntityProp, cmd chan int) *Entity {
	return &Entity{
		Prop:     prop,
		Movement: mv.NewMovement(prop.Movement()),
		SpaceId:  0,
		Command:  cmd,
	}
}
