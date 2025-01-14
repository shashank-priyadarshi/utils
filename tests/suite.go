package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"go.ssnk.in/utils/logger"
	loggerPorts "go.ssnk.in/utils/logger/ports"
	"go.ssnk.in/utils/tests/integration"
	"go.ssnk.in/utils/tests/load"
	"go.ssnk.in/utils/tests/profile"
	"go.ssnk.in/utils/tests/types"
	"gopkg.in/yaml.v3"
)

type Tester interface {
	Execute([]types.Config) error
}

const (
	FlagDescriptionConfigPath string = `Config for each supported test in a YAML.
	Default "configPath" value is "./config.yaml".`
	FlagDescriptionTests = `Runs Unit and Integration tests by default.
	To run all available tests in the suite, provide comma-separated values from 0 to 3, where
	1: Integration Tests,
	2: Integration Tests with Profiling enabled,
	3: Load Tests.
	If this flag is not provided through args or config.yaml, all integration tests are run.
	If both flag and config.yaml are detected, config.yaml takes precedence.`
	FlagDescriptionParallel = `If multiple tests are to be run in the suite, starts tests in parallel by default.
	Set "parallel" flag value to false to disable parallel execution.`
)

func main() {

	s := suite{
		log:   logger.New(logger.SetProvider("slog"), logger.SetLevel("info"), logger.SetFormat("json"), logger.WithTracing()),
		tests: make(map[types.Test]Tester),
	}

	s.builders = map[types.Test]func() *suite{
		types.Integration: s.withIntegrationTests,
		types.Profile:     s.withProfiling,
		types.Load:        s.withLoadTests,
	}

	var (
		configPath, testsStr string
		parallel             bool
		tests                []types.Test
	)

	flag.StringVar(&configPath, "config", "./config.yaml", FlagDescriptionConfigPath)
	flag.StringVar(&testsStr, "tests", "0", FlagDescriptionTests)
	flag.BoolVar(&parallel, "parallel", true, FlagDescriptionParallel)

	testsStrArr := strings.Split(testsStr, ",")
	for _, testStr := range testsStrArr {
		test, err := strconv.ParseInt(testStr, 10, 8)
		if err != nil {
			s.log.Warn(err.Error())
			continue
		}

		tests = append(tests, types.Test(test))
	}

	f, err := os.Open(configPath)
	if err != nil {
		s.log.Panic(err)
	}

	configBody, err := io.ReadAll(f)
	if err != nil {
		s.log.Panic(err)
	}

	var config []types.Config
	err = yaml.Unmarshal(configBody, &config)
	if err != nil {
		err = fmt.Errorf("invalid config passed: %w", err)
		s.log.Error(err)
		return
	}

	s.config = config

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
	builders map[types.Test]func() *suite
	tests    map[types.Test]Tester
	config   []types.Config
}

func (s *suite) withIntegrationTests() *suite {

	s.tests[types.Integration] = integration.New(s.log)
	return s
}

func (s *suite) withProfiling() *suite {

	s.tests[types.Profile] = profile.New(s.log)
	return s
}

func (s *suite) withLoadTests() *suite {

	s.tests[types.Load] = load.New(s.log)
	return s
}
