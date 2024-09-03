package mongodb

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/strikesecurity/strikememongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.ssnk.in/utils/types"
)

func Test_Create(t *testing.T) {
	// logger := ports.NewMockLogger(t)
	server, err := strikememongo.Start("4.0.5")
	defer server.Stop()
	assert.Equal(t, nil, err)

	opts := options.ClientOptions{}
	opts.ApplyURI(server.URI())
	client, err := mongo.Connect(context.TODO(), &opts)
	assert.Equal(t, nil, err)

	handle := Handle(client)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			documents := []interface{}{bson.D{{"foo", "bar"}}, bson.D{{"hello", "world"}}, bson.D{{"pi", 3.14159}}}
			_, err = handle.Create(context.TODO(), "test", "test", documents)

			assert.Equal(t, nil, err)
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Query(t *testing.T) {
	// logger := ports.NewMockLogger(t)
	server, err := strikememongo.Start("4.0.5")
	defer server.Stop()
	assert.Equal(t, nil, err)

	opts := options.ClientOptions{}
	opts.ApplyURI(server.URI())
	client, err := mongo.Connect(context.TODO(), &opts)
	assert.Equal(t, nil, err)

	handle := Handle(client)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			documents := []interface{}{bson.D{{"foo", "bar"}}, bson.D{{"hello", "world"}}, bson.D{{"pi", 3.14159}}}
			_, _ = handle.Create(context.TODO(), "test", "test", documents)

			response, err := handle.Query(context.TODO(), "test", "test", documents[1])
			assert.Equal(t, nil, err)

			values := response.Result[0].(bson.D)
			valueMap := values.Map()
			assert.Equal(t, "world", valueMap["hello"])
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Update(t *testing.T) {
	// logger := ports.NewMockLogger(t)
	server, err := strikememongo.Start("4.0.5")
	defer server.Stop()
	assert.Equal(t, nil, err)

	opts := options.ClientOptions{}
	opts.ApplyURI(server.URI())
	client, err := mongo.Connect(context.TODO(), &opts)
	assert.Equal(t, nil, err)

	handle := Handle(client)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			query := []interface{}{bson.D{{"foo", "bar"}}, bson.D{{"hello", "world"}}, bson.D{{"pi", 3.14159}}}
			updateQuery := bson.D{{"$set", bson.D{{"pi", "143"}}}}
			_, _ = handle.Create(context.TODO(), "test", "test", query)

			_, err = handle.Update(context.TODO(), "test", "test", query[2], updateQuery)
			assert.Equal(t, nil, err)

			response, err := handle.Query(context.TODO(), "test", "test", bson.D{{"pi", "143"}})
			assert.Equal(t, nil, err)

			values := response.Result[0].(bson.D)
			valueMap := values.Map()
			assert.Equal(t, "143", valueMap["pi"])
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Delete(t *testing.T) {
	server, err := strikememongo.Start("4.0.5")
	defer server.Stop()
	assert.Equal(t, nil, err)

	opts := options.ClientOptions{}
	opts.ApplyURI(server.URI())
	client, err := mongo.Connect(context.TODO(), &opts)
	assert.Equal(t, nil, err)

	handle := Handle(client)

	tests := []types.Test{{
		Name: "happy path",
		TestCase: func(t *testing.T) {
			query := []interface{}{bson.D{{"foo", "bar"}}}
			_, _ = handle.Create(context.TODO(), "test", "test", query)

			_, err = handle.Delete(context.TODO(), "test", "test", query[0])
			assert.Equal(t, nil, err)

			_, err = handle.Query(context.TODO(), "test", "test", bson.D{{"foo", "bar"}})
			assert.Equal(t, nil, err)

			// TODO: Result should contain deleted count
			// values := response.Result[0].(int64)
			// assert.Equal(t, 1, values)
		},
	}}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}
