package connections

import (
	"context"
	"fmt"

	redisclient "github.com/redis/go-redis/v9"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
)

func Redis(ctx context.Context, log ports.Logger, config *models.Config) (client *redisclient.Client, err error) {

	if len(config.Options.URI) == 0 {
		err = fmt.Errorf("redis uri cannot be empty")
		log.Error(err)
		return
	}

	// Other options pending: Using functional options
	opts := &redisclient.Options{
		Addr:     config.Options.URI,
		Username: config.Options.Username,
		Password: config.Options.Password,
	}

	client = redisclient.NewClient(opts)
	if client.Ping(ctx); err != nil {
		err = fmt.Errorf("error pinging redis: %v", err)
		log.Error(err)

		client.Close()
		return
	}

	return
}
