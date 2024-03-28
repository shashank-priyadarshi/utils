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

type Handler struct {
	log    ports.Logger
	client *mongo.Client
}

func Handle(log ports.Logger, client *mongo.Client) *Handler {
	return &Handler{
		log:    log,
		client: client,
	}
}

func (h *Handler) Create(ctx context.Context, args ...interface{}) (*models.Response, error) {

	if len(args) < 3 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err        error
		collection string
		database   string
		documents  []interface{}
		ok         bool
	)

	if database, ok = args[0].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "database")
	}

	if collection, ok = args[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "collection")
	}

	if documents, ok = args[2].([]interface{}); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "documents")
	}

	if _, err = h.client.Database(database).Collection(collection).InsertMany(ctx, documents); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error inserting entries to collection %s of database %s: %v", collection, database, err))
	}

	return nil, nil
}

func (h *Handler) Query(ctx context.Context, args ...interface{}) (*models.Response, error) {

	if len(args) < 3 {
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

	if database, ok = args[0].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "database")
	}

	if collection, ok = args[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "collection")
	}

	if query, ok = args[2].(bson.D); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "query")
	}

	if cur, err = h.client.Database(database).Collection(collection).Find(ctx, query); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error executing query %+v on collection %s in database %s", query, collection, database))
	}

	response := &models.Response{
		Result: make([]interface{}, 0),
	}

	for cur.Next(context.Background()) {
		var entry interface{}

		if err = cur.Decode(&entry); err != nil {
			return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error unmarshaling fetched document %+v to array of entries: %v", cur, err))
		}

		response.Result = append(response.Result, entry)
	}

	return response, nil
}

func (h *Handler) Update(ctx context.Context, args ...interface{}) (*models.Response, error) {

	if len(args) < 4 {
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

	if database, ok = args[0].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "database")
	}

	if collection, ok = args[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "collection")
	}

	if filterQuery, ok = args[2].(bson.D); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "filter query")
	}

	if updateQuery, ok = args[3].(bson.D); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "update query")
	}

	if updateResult, err = h.client.Database(database).Collection(collection).UpdateOne(ctx, filterQuery, updateQuery); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error updating entry in collection %s of database %s: %v", collection, database, err))
	}

	return &models.Response{Result: []interface{}{struct {
		MatchedCount, ModifiedCount int64
		UpsertedID                  interface{}
	}{
		MatchedCount:  updateResult.MatchedCount,
		ModifiedCount: updateResult.ModifiedCount,
		UpsertedID:    updateResult.UpsertedID,
	}}}, nil
}

func (h *Handler) Delete(ctx context.Context, args ...interface{}) (*models.Response, error) {

	if len(args) < 3 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err                  error
		collection, database string
		ok                   bool

		filterQuery  bson.D
		deleteResult *mongo.DeleteResult
	)

	if database, ok = args[0].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "database")
	}

	if collection, ok = args[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "collection")
	}

	if filterQuery, ok = args[2].(bson.D); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "filter query")
	}

	if deleteResult, err = h.client.Database(database).Collection(collection).DeleteOne(ctx, filterQuery); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error deleting entry from collection %s of database %s: %v", collection, database, err))
	}

	return &models.Response{Result: []interface{}{deleteResult.DeletedCount}}, nil
}

func (h *Handler) Begin(ctx context.Context, args ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Execute(ctx context.Context, args ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Rollback(ctx context.Context, args ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Configure(ctx context.Context, args ...interface{}) error {
	return nil
}

func (h *Handler) Close() error {
	return nil
}
