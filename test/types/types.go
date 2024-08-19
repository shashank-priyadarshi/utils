package types

type Config struct {
	Packages    []Package   `yaml:"packages"`
	Integration Integration `yaml:"integration"`
	Profile     Profile     `yaml:"profile"`
	Load        Load        `yaml:"load"`
}

type Package int

const (
	DATA Package = iota
	DATABASE
	LOGGER
	NETWORK
	PUBSUB
	SECURITY
	WORKER
)

type Integration struct {
}

type Profile struct {
}

type Load struct {
}
