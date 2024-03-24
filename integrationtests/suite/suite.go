package suite

import (
	"github.com/shashank-priyadarshi/utilities/logger/constants"
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
	"os"
)

type testSuite struct {
	log      loggerPort.Logger
	packages []string
}

type dependencies interface {
	fetch() string
}

func Test() {
	os.Setenv("LOG_PROVIDER", string(constants.SLOG))
	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("LOG_FORMAT", "JSON")
	os.Setenv("LOG_TRACE", "true")

	t := testSuite{
		packages: []string{"logger"},
	}

	for _, pkg := range t.packages {
		switch pkg {
		case "logger":
			t.testLogger(t.fetch())
		}
	}

}

func (t *testSuite) fetch() []string {
	return []string{os.Getenv("LOG_PROVIDER"), os.Getenv("LOG_LEVEL"), os.Getenv("LOG_FORMAT"), os.Getenv("LOG_TRACE")}
}
