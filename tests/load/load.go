package load

import (
	"go.ssnk.in/utils/tests/types"
)

type Load struct{}

func New() *Load {
	return &Load{}
}

func (l *Load) Execute(c *types.Config) error {
	return nil
}
