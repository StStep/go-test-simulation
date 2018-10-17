package unit

import (
	"github.com/StStep/go-test-simulation/internal/id"
)

type Unit interface {
	Id() id.Uid
}
