package utilities

import "testing"

type Test struct {
	Name     string
	TestCase func(*testing.T)
}
