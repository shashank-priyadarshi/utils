package types

import "fmt"

type Config struct {
	Package Package `json:"name"`
	Config  map[Test]struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	} `yaml:"config,omitempty"`
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type Alias Config
	aux := &struct {
		Package int `yaml:"name"`
		Config  map[int]struct {
			Key   string `yaml:"key"`
			Value string `yaml:"value"`
		} `yaml:"config,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := unmarshal(aux); err != nil {
		return err
	}

	c.Package = Package(aux.Package)
	c.Config = make(map[Test]struct {
		Key   string `yaml:"key"`
		Value string `yaml:"value"`
	})
	for k, v := range aux.Config {
		test := Test(k)
		if test < Integration || test > Load {
			return fmt.Errorf("invalid test value: %d", k)
		}
		c.Config[test] = v
	}

	return nil
}

type Package int

const (
	Algo Package = iota
	Database
	Logger
)

type Test int

const (
	Integration Test = iota
	Profile
	Load
)
