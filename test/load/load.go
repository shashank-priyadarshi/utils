package load

import (
	"go.ssnk.in/utils/test/types"
)

type Load struct{}

func New() *Load {
	return &Load{}
}

func (l *Load) Execute(c *types.Config) error {
	return nil
}
