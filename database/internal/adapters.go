package adapters

import (
	"context"
	"fmt"

	"github.com/shashank-priyadarshi/utilities/database/constants"
	"github.com/shashank-priyadarshi/utilities/database/internal/connections"
	"github.com/shashank-priyadarshi/utilities/database/internal/mongodb"
	"github.com/shashank-priyadarshi/utilities/database/internal/rdbms"
	"github.com/shashank-priyadarshi/utilities/database/internal/redis"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/database/ports"
	loggerport "github.com/shashank-priyadarshi/utilities/logger/ports"
)

func New(ctx context.Context, log loggerport.Logger, config *models.Config) (handle ports.Database, err error) {

	switch config.Type {
	case constants.MONGODB:
		client, err := connections.MongoDB(ctx, log, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to mongo db: %w", err)
		}

		handle = mongodb.Handle(log, client)

	case constants.MYSQLDB:
		client, err := connections.RDBMS(ctx, log, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to rdbms: %w", err)
		}

		handle, err = rdbms.Handle(log, config.Options.WithORM, config.Options.ORM, client)
		if err != nil {
			return nil, fmt.Errorf("error creating relational db handle: %w", err)
		}

	case constants.REDIS:
		client, err := connections.Redis(ctx, log, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to redis: %w", err)
		}

		handle = redis.Handle(log, client)

	default:
		return nil, fmt.Errorf("database type %s is not supported", config.Type)

	}

	return
}
