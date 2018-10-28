package configuration

import (
	phy "github.com/StStep/go-test-simulation/internal/physics/prop"
	"github.com/StStep/go-test-simulation/internal/unit"
	form "github.com/StStep/go-test-simulation/internal/unit/formation"
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
	config.Entities["swordman"] = &unit.EntProp{"swordman", pConf}

	config.Formations[""] = &form.Prop{"", 5, 0.5, 0.5}

	membs := make(map[string]int)
	membs["swordman"] = 20
	config.Units["swords"] = &unit.UnitProp{"swords", membs, []string{""}}

	tables := []struct {
		wpath  string
		wpanic bool
		rpath  string
		rpanic bool
	}{
		{"test.json", false, "test.json", false},
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
