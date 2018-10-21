package concrete

import (
	conf "github.com/StStep/go-test-simulation/internal/configuration"
	"github.com/StStep/go-test-simulation/internal/entity"
	pr "github.com/StStep/go-test-simulation/internal/entity/prop"
	"github.com/StStep/go-test-simulation/internal/id"
	"github.com/StStep/go-test-simulation/internal/ledger"
	"github.com/StStep/go-test-simulation/internal/physics"
	fl "gonum.org/v1/gonum/floats"
)

type constructor struct {
	db      ledger.LedgerRO
	conf    conf.Configuration
	idgen   id.EidGen
	physics physics.Physics
}

type concrete struct {
	db         ledger.LedgerRO
	prop       pr.Prop
	id         id.Eid
	unit       id.Uid
	physics    physics.Physics
	command    chan int
	formOffset [2]float64
}

func NewConstructor(db ledger.LedgerRO, conf conf.Configuration, idgen id.EidGen, physics physics.Physics) entity.EntityConstructor {
	return &constructor{db: db, conf: conf, idgen: idgen, physics: physics}
}

func (c *constructor) New(name string, uid id.Uid, cmd chan int, pos [2]float64, offset [2]float64) entity.Entity {
	prop := c.conf.Entity(name)
	e := concrete{
		db:         c.db,
		prop:       prop,
		id:         c.idgen.Id(),
		unit:       uid,
		physics:    c.physics,
		command:    cmd,
		formOffset: offset,
	}
	c.physics.RegisterEntity(e.id, prop.Physics(), pos)
	return &e
}

func (e *concrete) Prop() pr.Prop {
	return e.prop
}

func (e *concrete) Id() id.Eid {
	return e.id
}

func (e *concrete) Position() [2]float64 {
	return e.physics.Position(e.Id())
}

func (e *concrete) Velocity() [2]float64 {
	return e.physics.Velocity(e.Id())
}

func (e *concrete) FormError() (dir [2]float64, dist float64) {
	g := e.db.Unit(e.unit).Guide()
	pos := e.physics.Position(e.Id())
	fl.Sub(g[:], pos[:])
	dist = fl.Norm(g[:], 2)
	dir = g
	fl.Scale(1/dist, dir[:])
	return
}

func (e *concrete) LogicStep(del float64) {
	// TODO Currently only setting CmdVel depending upon offset and pos
	dir, dist := e.FormError()

	// TODO Shouldn't use dist as speed, just rough placeholder
	e.physics.SetCommand(e.Id(), dir, dist)
}
