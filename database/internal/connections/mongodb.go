package connections

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.ssnk.in/utils/database/models"
)

var opts = &options.ClientOptions{}

func MongoDB(ctx context.Context, config *models.Config) (client *mongo.Client, err error) {

	if len(config.Options.URI) == 0 {
		err = fmt.Errorf("mongo db uri cannot be empty")
		log.Error(err)
		return
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
		NilSliceAsEmpty:   true,
	}

	opts.ApplyURI(config.Options.URI).SetServerAPIOptions(serverAPI).SetMaxPoolSize(uint64(config.Options.DatabaseOptions.MongoDBPoolSize)).SetBSONOptions(bsonOpts)

	// TODO: Create collection if not exist
	/*dbClient := client.Database(LINCLIST_DATABASE)

	existingCollections, err := dbClient.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error fetching existing collection name list: %v", err)
	}

	var existingCollectionMap = make(map[string]any)
	for _, collection := range existingCollections {
		existingCollectionMap[collection] = nil
	}

	for _, collection := range config.Collections {
		if _, ok := existingCollectionMap[collection]; !ok {
			err := dbClient.CreateCollection(context.TODO(), collection)
			if err != nil {
				fmt.Println("error while creating collection", collection, ":", err)
				return nil, fmt.Errorf("error creating collection %s: %v", collection, err)
			}
			indexes, err := dbClient.Collection(collection).Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
				{Keys: bson.D{{"tags", []string{}}}, Options: options.Index().SetUnique(false)},
				{Keys: bson.D{{"metadata.read", false}}, Options: options.Index().SetUnique(false)},
				{Keys: bson.D{{"metadata.stale", false}}, Options: options.Index().SetUnique(false)},
			})

			if err != nil {
				fmt.Println("error while creating indexes in collection", collection, ":", err)
				return nil, fmt.Errorf("error creating indexes for collection %s: %v", collection, err)
			}

			fmt.Println("newly created indexes on collection ", collection, ":", indexes)
		}
	}*/

	// TODO: Backoff connection logic
	if client, err = mongo.Connect(ctx, opts); err != nil {
		err = fmt.Errorf("error connecting to database: %v", err)
		log.Error(err)
		return
	}

	if err = client.Ping(ctx, nil); err != nil {
		err = fmt.Errorf("error pinging mongo db on established connection: %v", err)
		log.Error(err)
		// TODO: Backoff disconnection logic
		client.Disconnect(ctx)
		return
	}

	return
}
