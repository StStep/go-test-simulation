package concrete

import (
	conf "github.com/StStep/go-test-simulation/internal/configuration"
	"github.com/StStep/go-test-simulation/internal/entity"
	pr "github.com/StStep/go-test-simulation/internal/entity/prop"
	"github.com/StStep/go-test-simulation/internal/id"
	"github.com/StStep/go-test-simulation/internal/ledger"
	mv "github.com/StStep/go-test-simulation/internal/movement"
	"github.com/StStep/go-test-simulation/internal/space"
)

type constructor struct {
	db    ledger.LedgerRO
	conf  conf.Configuration
	idgen id.EidGen
	space space.Space
}

type concrete struct {
	prop       pr.Prop
	id         id.Eid
	movement   *mv.Movement
	space      space.Space
	command    chan int
	formOffset [2]float64
}

func NewConstructor(db ledger.LedgerRO, conf conf.Configuration, idgen id.EidGen, space space.Space) entity.EntityConstructor {
	return &constructor{db: db, conf: conf, idgen: idgen, space: space}
}

func (c *constructor) New(name string, cmd chan int, pos [2]float64, offset [2]float64) entity.Entity {
	prop := c.conf.Entity(name)
	e := concrete{
		prop:       prop,
		id:         c.idgen.Id(),
		movement:   mv.NewMovement(prop.Movement()),
		space:      c.space,
		command:    cmd,
		formOffset: offset,
	}
	c.space.RegisterEntity(e.id, pos, prop.Radius())
	return &e
}

func (e *concrete) Prop() pr.Prop {
	return e.prop
}

func (e *concrete) Id() id.Eid {
	return e.id
}

// TODO Use Updater and Viewer in update step to follow CmdReform
