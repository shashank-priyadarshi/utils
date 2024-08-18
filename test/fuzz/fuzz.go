package fuzz

import (
	"github.com/shashank-priyadarshi/utilities/test/types"
)

type Fuzz struct{}

func New() *Fuzz {
	return &Fuzz{}
}

func (i *Fuzz) Execute(c *types.Config) error {
	return nil
}
