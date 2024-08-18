package mutation

import (
	"github.com/shashank-priyadarshi/utilities/test/types"
)

type Mutation struct{}

func New() *Mutation {
	return &Mutation{}
}

func (i *Mutation) Execute(c *types.Config) error {
	return nil
}
