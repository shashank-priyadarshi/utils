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
	ports2 "github.com/shashank-priyadarshi/utilities/logger/ports"
)

func NewDatabaseAdapter(ctx context.Context, log ports2.Logger, config *models.Config) (handle ports.Database, err error) {

	switch config.Type {
	case constants.MONGODB:
		client, err := connections.NewMongoDBClient(ctx, log, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to mongo db: %w", err)
		}

		handle = mongodb.NewMongoDBHandle(log, client)

	case constants.MYSQLDB:
		client, err := connections.NewRDBMSClient(ctx, log, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to rdbms: %w", err)
		}

		handle, err = rdbms.NewRelationalDBHandle(log, config.Options.WithORM, config.Options.ORM, client)
		if err != nil {
			return nil, fmt.Errorf("error creating relational db handle: %w", err)
		}

	case constants.REDIS:
		client, err := connections.NewRedisClient(ctx, log, config)
		if err != nil {
			return nil, fmt.Errorf("error connecting to redis: %w", err)
		}

		handle = redis.NewRedisHandle(log, client)

	default:
		return nil, fmt.Errorf("database type %s  is not supported", config.Type)

	}

	return
}
