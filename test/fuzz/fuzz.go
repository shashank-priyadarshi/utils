package fuzz

import (
	"github.com/shashank-priyadarshi/utilities/test/types"
)

// https://go.dev/doc/security/fuzz/

type Fuzz struct{}

func New() *Fuzz {
	return &Fuzz{}
}

func (i *Fuzz) Execute(c *types.Config) error {
	return nil
}
