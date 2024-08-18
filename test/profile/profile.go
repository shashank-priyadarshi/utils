package profile

import (
	"github.com/shashank-priyadarshi/utilities/test/types"
)

type Profile struct{}

func New() *Profile {
	return &Profile{}
}

func (i *Profile) Execute(c *types.Config) error {
	return nil
}
