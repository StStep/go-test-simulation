package configuration

import (
	"encoding/json"
	ent "github.com/StStep/go-test-simulation/internal/entity"
	form "github.com/StStep/go-test-simulation/internal/formation"
	un "github.com/StStep/go-test-simulation/internal/unit"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Entities   map[string]*ent.Prop  `json:"entities"`
	Units      map[string]*un.Prop   `json:"units"`
	Formations map[string]*form.Prop `json:"formations"`
}

func New() *Configuration {
	c := Configuration{}
	c.Entities = make(map[string]*ent.Prop)
	c.Units = make(map[string]*un.Prop)
	c.Formations = make(map[string]*form.Prop)
	return &c
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func FromFile(path string) *Configuration {
	j, err := ioutil.ReadFile(path)
	check(err)

	c := Configuration{}
	err = json.Unmarshal(j, &c)
	check(err)

	return &c
}

func (c *Configuration) ToFile(path string) {
	j, err := json.Marshal(c)
	check(err)

	f, err := os.Create(path)
	check(err)
	defer f.Close()
	f.Write(j)
}
