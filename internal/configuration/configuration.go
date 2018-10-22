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
	c.entities = make(map[string]ent.Prop)
	c.units = make(map[string]un.Prop)
	c.formations = make(map[string]form.FormationProp)
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
	entities   map[string]ent.Prop
	units      map[string]un.Prop
	formations map[string]form.FormationProp
}

func (c *conf) Entity(name string) ent.Prop {
	return c.entities[name]
}

func (c *conf) Unit(name string) un.Prop {
	return c.units[name]
}

func (c *conf) Formation(name string) form.FormationProp {
	return c.formations[name]
}

func (c *conf) ToFile(path string) {
	j, err := json.Marshal(c)
	check(err)

	f, err := os.Create(path)
	check(err)
	defer f.Close()
	f.Write(j)
}
