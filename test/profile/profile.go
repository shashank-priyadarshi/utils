package profile

import (
	"go.ssnk.in/utils/test/types"
)

type Profile struct{}

func New() *Profile {
	return &Profile{}
}

func (m *Profile) Execute(c *types.Config) error {
	return nil
}
