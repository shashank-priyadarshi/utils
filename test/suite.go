package main

import (
	"flag"
	"go.ssnk.in/utils/logger"
	loggerPorts "go.ssnk.in/utils/logger/ports"
	"go.ssnk.in/utils/test/integration"
	"go.ssnk.in/utils/test/load"
	"go.ssnk.in/utils/test/profile"
	"go.ssnk.in/utils/test/types"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Tester interface {
	Execute(*types.Config) error
}

type Test int

const (
	Integration Test = iota
	Profile
	Load
)

const (
	FlagDescriptionConfigPath string = `Config for each supported test in a YAML.
	Default "configPath" value is "./config.yaml".`
	FlagDescriptionTests = `Runs Unit and Integration tests by default.
	To run all available tests in the suite, provide comma-separated values from 0 to 5, where
	1: Integration Tests,
	2: Integration Tests with Profiling enabled,
	3: Load Tests.`
	FlagDescriptionParallel = `If multiple tests are to be run in the suite, starts tests in parallel by default.
	Set "parallel" flag value to false to disable parallel execution.`
)

func main() {

	s := suite{
		log: logger.New(logger.SetLevel("info"), logger.SetProvider("slog")),
	}

	s.builders = map[Test]func() *suite{
		Integration: s.withIntegrationTests,
		Profile:     s.withProfiling,
		Load:        s.withLoadTests,
	}

	var (
		configPath, testsStr string
		parallel             bool
		tests                []Test
	)

	flag.StringVar(&configPath, "config", "./config.yaml", FlagDescriptionConfigPath)
	flag.StringVar(&testsStr, "tests", "0,3", FlagDescriptionTests)
	flag.BoolVar(&parallel, "parallel", true, FlagDescriptionParallel)

	testsStrArr := strings.Split(testsStr, ",")
	for _, testStr := range testsStrArr {
		test, err := strconv.ParseInt(testStr, 10, 8)
		if err != nil {
			s.log.Warn(err.Error())
			continue
		}

		tests = append(tests, Test(test))
	}

	f, err := os.Open(configPath)
	if err != nil {
		s.log.Panic(err)
	}

	configBody, err := io.ReadAll(f)
	if err != nil {
		s.log.Panic(err)
	}

	var config types.Config
	err = yaml.Unmarshal(configBody, &config)
	if err != nil {
		s.log.Panic(err)
	}

	s.config = &config

	for _, t := range tests {
		s.builders[t]()
	}

	switch parallel {
	case false:
		s.startExecution()
	default:
		s.startParallelExecution()
	}
}

func (s *suite) startParallelExecution() {
	var wg sync.WaitGroup
	wg.Add(len(s.tests))

	for _, t := range s.tests {
		go func(t Tester) {
			defer wg.Done()
			if t == nil {
				return
			}

			if err := t.Execute(s.config); err != nil {
				s.log.Error(err)
			}
		}(t)
	}

	wg.Wait()
}

func (s *suite) startExecution() {
	for _, t := range s.tests {
		if t == nil {
			continue
		}
		if err := t.Execute(s.config); err != nil {
			s.log.Error(err)
		}
	}
}

type suite struct {
	log      loggerPorts.Logger
	builders map[Test]func() *suite
	tests    map[Test]Tester
	config   *types.Config
}

func (s *suite) withIntegrationTests() *suite {

	s.tests[Integration] = integration.New()
	return s
}

func (s *suite) withProfiling() *suite {

	s.tests[Profile] = profile.New()
	return s
}

func (s *suite) withLoadTests() *suite {

	s.tests[Load] = load.New()
	return s
}
