package types

type Job func() []interface{}

type Status string

func (s *Status) String() string {
	return string(*s)
}
