package entity

import (
	"github.com/StStep/go-test-simulation/internal/id"
	phy "github.com/StStep/go-test-simulation/internal/physics/prop"
)

type Entity struct {
	Id         id.Eid
	Prop       *Prop
	Unit       id.Uid
	Command    chan int
	FormOffset [2]float64
}

// TODO Could decorate props with modifiers? Or decorate with a sturct allowing modifiers?
type Prop struct {
	Name    string    `json:"name"`
	Physics *phy.Prop `json:"physics"`
}
