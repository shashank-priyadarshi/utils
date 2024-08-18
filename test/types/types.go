package types

type Config struct {
	Packages    []Package   `yaml:"packages"`
	Unit        Unit        `yaml:"unit"`
	Mutation    Mutation    `yaml:"mutation"`
	Fuzz        Fuzz        `yaml:"fuzz"`
	Integration Integration `yaml:"integration"`
	Profile     Profile     `yaml:"profile"`
	Benchmark   Benchmark   `yaml:"benchmark"`
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

type Unit struct {
}

type Mutation struct {
}

type Fuzz struct {
}

type Integration struct {
}

type Profile struct {
}

type Benchmark struct {
}
