package types

import "context"

type Job func(ctx context.Context) []interface{}

type Status string

func (s *Status) String() string {
	return string(*s)
}
