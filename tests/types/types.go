package types

type Config struct {
	Packages    []Package   `yaml:"packages"`
	Integration Integration `yaml:"integration"`
	Profile     Profile     `yaml:"profile"`
	Load        Load        `yaml:"load"`
}

type Package int

const (
	Algo Package = iota
	Database
	Logger
)

type Integration struct{}

type Profile struct{}

type Load struct{}
