package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.ssnk.in/utils/types"
)

func Test_Create(t *testing.T) {
	// logger := ports.NewMockLogger(t)
	miniredisServer := miniredis.NewMiniRedis()
	if err := miniredisServer.Start(); err != nil {
		t.Fatal("failed to start redis test server using mini redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr: miniredisServer.Addr(),
	})

	handle := Handle(client)

	tests := []types.Test{
		{
			Name: "happy path",
			TestCase: func(t *testing.T) {
				key := "1"
				value := key
				result, err := handle.Create(context.TODO(), key, value)

				assert.Equal(t, nil, err)
				assert.Equal(t, "OK", result.Result[0])
			},
		},
	}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Query(t *testing.T) {
	// logger := ports.NewMockLogger(t)
	miniredisServer := miniredis.NewMiniRedis()
	if err := miniredisServer.Start(); err != nil {
		t.Fatal("failed to start redis test server using mini redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr: miniredisServer.Addr(),
	})

	handle := Handle(client)

	tests := []types.Test{
		{
			Name: "happy path",
			TestCase: func(t *testing.T) {
				key := "1"
				value := key
				_, _ = handle.Create(context.TODO(), key, value)

				result, err := handle.Query(context.TODO(), key)

				assert.Equal(t, nil, err)
				assert.Equal(t, value, result.Result[0].(string))
			},
		},
	}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Update(t *testing.T) {
	// logger := ports.NewMockLogger(t)
	miniredisServer := miniredis.NewMiniRedis()
	if err := miniredisServer.Start(); err != nil {
		t.Fatal("failed to start redis test server using mini redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr: miniredisServer.Addr(),
	})

	handle := Handle(client)

	tests := []types.Test{
		{
			Name: "happy path",
			TestCase: func(t *testing.T) {
				key := "1"
				value := key
				_, _ = handle.Create(context.TODO(), key, value)

				value += value
				result, err := handle.Update(context.TODO(), key, value)

				assert.Equal(t, nil, err)
				assert.Equal(t, "OK", result.Result[0].(string))

				result, err = handle.Query(context.TODO(), key)

				assert.Equal(t, nil, err)
				assert.Equal(t, value, result.Result[0].(string))
			},
		},
	}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}

func Test_Delete(t *testing.T) {
	// logger := ports.NewMockLogger(t)
	miniredisServer := miniredis.NewMiniRedis()
	if err := miniredisServer.Start(); err != nil {
		t.Fatal("failed to start redis test server using mini redis")
	}

	client := redis.NewClient(&redis.Options{
		Addr: miniredisServer.Addr(),
	})

	handle := Handle(client)

	tests := []types.Test{
		{
			Name: "happy path",
			TestCase: func(t *testing.T) {
				key := "1"

				result, err := handle.Delete(context.TODO(), key)

				assert.Equal(t, nil, err)
				assert.Equal(t, int64(0), result.Result[0].(int64))

				result, err = handle.Query(context.TODO(), key)

				assert.Equal(t, ": : operation failed: redis: nil", err.Error())
			},
		},
	}

	for id, test := range tests {
		t.Run(fmt.Sprintf("Test ID: %d, Test Name: %s", id, test.Name), test.TestCase)
	}
}
