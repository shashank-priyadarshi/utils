package mongodb

import (
	"context"
	"fmt"

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

func (h *Handle) Create(ctx context.Context, params ...interface{}) (response *models.Response, err error) {
	if len(params) < 3 {
		err = fmt.Errorf("insufficient parameters")
		response.Error = err
		return
	}

	var (
		collection string
		database   string
		documents  []interface{}
		ok         bool
	)

	if database, ok = params[0].(string); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if collection, ok = params[1].(string); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if documents, ok = params[2].([]interface{}); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if _, err = h.client.Database(database).Collection(collection).InsertMany(context.TODO(), documents); err != nil {
		err = fmt.Errorf("error inserting entries to collection %s of database %s: %v", collection, database, err)
		response.Error = err
		return
	}

	return
}

func (h *Handle) Query(ctx context.Context, params ...interface{}) (response *models.Response, err error) {

	if len(params) < 3 {
		err = fmt.Errorf("insufficient parameters")
		response.Error = err
		return
	}

	var (
		collection string
		cur        *mongo.Cursor
		database   string
		query      bson.D
		ok         bool
	)

	if database, ok = params[0].(string); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if collection, ok = params[1].(string); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if query, ok = params[2].(bson.D); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if cur, err = h.client.Database(database).Collection(collection).Find(context.TODO(), query); err != nil {
		err = fmt.Errorf("error executing query %+v on collection %s in database %s", query, collection, database)
		response.Error = err
		return
	}

	for cur.Next(context.Background()) {
		var entry interface{}

		if err = cur.Decode(&entry); err != nil {
			err = fmt.Errorf("error unmarshaling fetched document %+v to array of entries: %v", cur, err)
			response.Error = err
			return
		}

		response.Result = append(response.Result, entry)
	}

	return
}

func (h *Handle) Update(ctx context.Context, params ...interface{}) (response *models.Response, err error) {
	if len(params) < 4 {
		err = fmt.Errorf("insufficient parameters")
		response.Error = err
		return
	}

	var (
		collection   string
		database     string
		filterQuery  bson.D
		updateQuery  bson.D
		updateResult *mongo.UpdateResult
		ok           bool
	)

	if database, ok = params[0].(string); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if collection, ok = params[1].(string); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if filterQuery, ok = params[2].(bson.D); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if updateQuery, ok = params[2].(bson.D); !ok {
		err = fmt.Errorf("invalid parameter")
		response.Error = err
		return
	}

	if updateResult, err = h.client.Database(database).Collection(collection).UpdateOne(context.TODO(), filterQuery, updateQuery); err != nil {
		err = fmt.Errorf("error updating entry in collection %s of database %s: %v", collection, database, err)
		response.Error = err
		return
	}

	response.Result = []interface{}{struct {
		MatchedCount, ModifiedCount int64
		UpsertedID                  interface{}
	}{
		MatchedCount:  updateResult.MatchedCount,
		ModifiedCount: updateResult.ModifiedCount,
		UpsertedID:    updateResult.UpsertedID,
	}}

	return
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
