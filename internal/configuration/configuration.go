package configuration

import (
	"encoding/json"
	ent "github.com/StStep/go-test-simulation/internal/entity/prop"
	form "github.com/StStep/go-test-simulation/internal/formationprop"
	un "github.com/StStep/go-test-simulation/internal/unit/prop"
	"io/ioutil"
	"os"
)

type Configuration interface {
	Entity(name string) ent.Prop
	Unit(name string) un.Prop
	Formation(name string) form.FormationProp
	ToFile(path string)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func newConf() *conf {
	c := conf{}
	c.Entities = make(map[string]*ent.Pprop)
	c.Units = make(map[string]*un.Pprop)
	c.Formations = make(map[string]*form.Pprop)
	return &c
}

func FromFile(path string) Configuration {
	j, err := ioutil.ReadFile(path)
	check(err)

	c := conf{}
	err = json.Unmarshal(j, &c)
	check(err)

	return &c
}

type conf struct {
	Entities   map[string]*ent.Pprop  `json:"entities"`
	Units      map[string]*un.Pprop   `json:"units"`
	Formations map[string]*form.Pprop `json:"formations"`
}

func (c *conf) Entity(name string) ent.Prop {
	return c.Entities[name]
}

func (c *conf) Unit(name string) un.Prop {
	return c.Units[name]
}

func (c *conf) Formation(name string) form.FormationProp {
	return c.Formations[name]
}

func (c *conf) ToFile(path string) {
	j, err := json.Marshal(c)
	check(err)

	f, err := os.Create(path)
	check(err)
	defer f.Close()
	f.Write(j)
}
