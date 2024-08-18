package benchmark

import (
	"github.com/shashank-priyadarshi/utilities/test/types"
)

type Benchmark struct{}

func New() *Benchmark {
	return &Benchmark{}
}

func (i *Benchmark) Execute(c *types.Config) error {
	return nil
}
