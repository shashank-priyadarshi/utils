package connections

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"

	redisclient "github.com/redis/go-redis/v9"
	"github.com/shashank-priyadarshi/utilities/database/models"
)

func Redis(ctx context.Context, config *models.Config) (client *redisclient.Client, err error) {

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

		client.Close()
		return
	}

	return
}
