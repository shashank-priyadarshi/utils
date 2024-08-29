package adapters

import (
	"context"
	"fmt"

	"go.ssnk.in/utils/database/constants"
	"go.ssnk.in/utils/database/internal/connections"
	"go.ssnk.in/utils/database/internal/mongodb"
	"go.ssnk.in/utils/database/internal/rdbms"
	"go.ssnk.in/utils/database/internal/redis"
	"go.ssnk.in/utils/database/models"
	"go.ssnk.in/utils/database/ports"
)

func New(ctx context.Context, config *models.Config) (handle ports.Database, err error) {

	switch config.Type {
	case constants.MONGODB:
		client, err := connections.MongoDB(ctx, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to mongo db: %w", err)
		}

		handle = mongodb.Handle(client)

	case constants.MYSQLDB:
		client, err := connections.RDBMS(ctx, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to rdbms: %w", err)
		}

		handle, err = rdbms.Handle(config.Options.WithORM, config.Options.ORM, client)
		if err != nil {
			return nil, fmt.Errorf("error creating relational db handle: %w", err)
		}

	case constants.REDIS:
		client, err := connections.Redis(ctx, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to redis: %w", err)
		}

		handle = redis.Handle(client)

	default:
		return nil, fmt.Errorf("database type %s is not supported", config.Type)

	}

	return
}
