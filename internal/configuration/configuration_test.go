package configuration

import (
	ent "github.com/StStep/go-test-simulation/internal/entity"
	form "github.com/StStep/go-test-simulation/internal/formation"
	phy "github.com/StStep/go-test-simulation/internal/physics/prop"
	un "github.com/StStep/go-test-simulation/internal/unit"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile(t *testing.T) {
	assert := assert.New(t)

	config := New()

	pConf := &phy.Prop{
		[4]float64{5, 5, 5, 5},
		[4]float64{5, 5, 5, 5},
		[4]float64{5, 5, 5, 5},
		[4]float64{},
		[4]float64{},
		0, 0,
		0.25}
	config.Entities["swordman"] = &ent.Prop{"swordman", pConf}

	config.Formations[""] = &form.Prop{"", 5, 0.5, 0.5}

	membs := make(map[string]int)
	membs["swordman"] = 20
	config.Units["swords"] = &un.Prop{"swords", membs, []string{""}}

	tables := []struct {
		wpath  string
		wpanic bool
		rpath  string
		rpanic bool
	}{
		{"config.json", false, "config.json", false},
	}

	for i, v := range tables {
		if v.wpanic {
			assert.Panicsf(func() { config.ToFile(v.wpath) }, "Test %v: No panic", i)
		} else {
			assert.NotPanicsf(func() { config.ToFile(v.wpath) }, "Test %v: Panic", i)
		}

		var n *Configuration
		if v.rpanic {
			assert.Panicsf(func() { FromFile(v.rpath) }, "Test %v: No panic", i)
		} else {
			assert.NotPanicsf(func() { n = FromFile(v.rpath) }, "Test %v: Panic", i)
		}

		if !v.wpanic && !v.rpanic {
			assert.Equalf(config, n, "Test %v: Not Equal", i)
		}
	}
}
