package connections

import (
	"context"
	"fmt"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
)

var opts *options.ClientOptions

func NewMongoDBClient(ctx context.Context, log ports.Logger, config *models.Config) (client *mongo.Client, err error) {

	if len(config.Options.URI) == 0 {
		err = fmt.Errorf("mongo db uri cannot be empty")
		log.Error(err)
		return
	}

	opts.ApplyURI(config.Options.URI)

	// Backoff connection logic
	if client, err = mongo.Connect(ctx, opts); err != nil {
		err = fmt.Errorf("error connecting to database: %v", err)
		log.Error(err)
		return
	}

	if err = client.Ping(ctx, nil); err != nil {
		err = fmt.Errorf("error pinging mongo db on established connection: %v", err)
		log.Error(err)
		// Backoff disconnection logic
		client.Disconnect(ctx)
		return
	}

	return
}
