package unit

import (
	"github.com/shashank-priyadarshi/utilities/test/types"
)

type Unit struct{}

func New() *Unit {
	return &Unit{}
}

func (i *Unit) Execute(c *types.Config) error {
	return nil
}
