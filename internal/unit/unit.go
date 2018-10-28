package unit

import (
	phy "github.com/StStep/go-test-simulation/internal/physics/prop"
)

type Unit struct {
	Id        uint64
	Prop      UnitProp
	GuideId   uint64
	MemberIds []uint64
}

func (u *Unit) Size() int {
	return len(u.MemberIds)
}

type UnitProp struct {
	Name       string         `json:"name"`
	Members    map[string]int `json:"members"`
	Formations []string       `json:"formations"`
}

func (p *UnitProp) Size() int {
	ret := 0
	for _, v := range p.Members {
		ret = ret + v
	}
	return ret
}

type Entity struct {
	Id         uint64
	Prop       *EntProp
	UnitId     uint64
	Command    chan int
	FormOffset [2]float64
}

// TODO Could decorate props with modifiers? Or decorate with a sturct allowing modifiers?
type EntProp struct {
	Name    string    `json:"name"`
	Physics *phy.Prop `json:"physics"`
}
