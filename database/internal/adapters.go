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
