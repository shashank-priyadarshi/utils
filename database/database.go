package database

import (
	"context"
	"fmt"

	"github.com/shashank-priyadarshi/utilities/database/constants"
	adapters "github.com/shashank-priyadarshi/utilities/database/internal"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/database/ports"
	"github.com/shashank-priyadarshi/utilities/logger"
	loggerport "github.com/shashank-priyadarshi/utilities/logger/ports"
)

func New(ctx context.Context, log loggerport.Logger, config models.Config) (database ports.Database, err error) {

	if !isSupported(config.Type) {
		err = fmt.Errorf("unsupported database type: %s", config.Type)
		return
	}

	if log == nil {
		err = fmt.Errorf("received uninitialized logger, initialized module logger")

		logOptions := config.Options.LogOptions
		log, err = logger.NewLogger(logOptions.Provider, logOptions.Level, logOptions.Format, logOptions.Trace)
		log.Warn("%v", err)
		log.With("module", "database")

		if err != nil {
			err = fmt.Errorf("error initializing new logger: %v", err)
			return
		}
	}

	return adapters.New(ctx, log, &config)
}

func isSupported(db constants.Database) bool {

	var supported = make(map[constants.Database]any)
	supported[constants.MYSQLDB] = nil
	supported[constants.MONGODB] = nil
	supported[constants.REDIS] = nil

	if _, ok := supported[db]; !ok {
		return false
	}

	return true
}
