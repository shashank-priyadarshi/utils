package types

import "testing"

type Test struct {
	Name     string
	TestCase func(t *testing.T)
}
