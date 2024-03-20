package mongodb

import (
	"context"
	"fmt"
	"github.com/shashank-priyadarshi/utilities"

	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handle struct {
	log    ports.Logger
	client *mongo.Client
}

func NewMongoDBHandle(log ports.Logger, client *mongo.Client) *Handle {
	return &Handle{
		log:    log,
		client: client,
	}
}

func (h *Handle) Create(ctx context.Context, params ...interface{}) (*models.Response, error) {

	if len(params) < 3 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err        error
		collection string
		database   string
		documents  []interface{}
		ok         bool
	)

	if database, ok = params[0].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "database")
	}

	if collection, ok = params[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "collection")
	}

	if documents, ok = params[2].([]interface{}); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "documents")
	}

	if _, err = h.client.Database(database).Collection(collection).InsertMany(context.TODO(), documents); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error inserting entries to collection %s of database %s: %v", collection, database, err))
	}

	return nil, nil
}

func (h *Handle) Query(ctx context.Context, params ...interface{}) (*models.Response, error) {

	if len(params) < 3 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err        error
		collection string
		cur        *mongo.Cursor
		database   string
		query      bson.D
		ok         bool
	)

	if database, ok = params[0].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "database")
	}

	if collection, ok = params[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "collection")
	}

	if query, ok = params[2].(bson.D); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "query")
	}

	if cur, err = h.client.Database(database).Collection(collection).Find(context.TODO(), query); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error executing query %+v on collection %s in database %s", query, collection, database))
	}

	var response *models.Response

	for cur.Next(context.Background()) {
		var entry interface{}

		if err = cur.Decode(&entry); err != nil {
			return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error unmarshaling fetched document %+v to array of entries: %v", cur, err))
		}

		response.Result = append(response.Result, entry)
	}

	return response, nil
}

func (h *Handle) Update(ctx context.Context, params ...interface{}) (*models.Response, error) {

	if len(params) < 4 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err          error
		collection   string
		database     string
		filterQuery  bson.D
		updateQuery  bson.D
		updateResult *mongo.UpdateResult
		ok           bool
	)

	if database, ok = params[0].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "database")
	}

	if collection, ok = params[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "collection")
	}

	if filterQuery, ok = params[2].(bson.D); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "filter query")
	}

	if updateQuery, ok = params[2].(bson.D); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "update query")
	}

	if updateResult, err = h.client.Database(database).Collection(collection).UpdateOne(context.TODO(), filterQuery, updateQuery); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error updating entry in collection %s of database %s: %v", collection, database, err))
	}

	var response *models.Response
	{
	}

	response.Result = []interface{}{struct {
		MatchedCount, ModifiedCount int64
		UpsertedID                  interface{}
	}{
		MatchedCount:  updateResult.MatchedCount,
		ModifiedCount: updateResult.ModifiedCount,
		UpsertedID:    updateResult.UpsertedID,
	}}

	return response, nil
}

func (h *Handle) Delete(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handle) Begin(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handle) Execute(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handle) Rollback(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handle) Configure(ctx context.Context, i ...interface{}) error {
	return nil
}

func (h *Handle) Close() error {
	return nil
}
