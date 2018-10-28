package logic

import (
	"github.com/StStep/go-test-simulation/internal/state"
	"github.com/StStep/go-test-simulation/internal/unit"
	fl "gonum.org/v1/gonum/floats"
)

func Error(e *unit.Entity, state *state.State) (dir [2]float64, dist float64) {
	g := state.Physics.Position(state.Ledger.UnitData[e.UnitId].GuideId)
	pos := state.Physics.Position(e.Id)
	fl.Sub(g[:], pos[:])
	dist = fl.Norm(g[:], 2)
	dir = g
	fl.Scale(1/dist, dir[:])
	return
}

func LogicStep(u *unit.Unit, state *state.State, del float64) {
	for _, v := range u.MemberIds {
		e := state.Ledger.EntityData[v]

		// TODO Currently only setting CmdVel depending upon offset and pos
		dir, dist := Error(e, state)

		// TODO Shouldn't use dist as speed, just rough placeholder
		state.Physics.SetCommand(e.Id, dir, dist)
	}
}
