package redis

import (
	"context"
	"github.com/shashank-priyadarshi/utilities"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Handle struct {
	log    ports.Logger
	client *redis.Client
}

func NewRedisHandle(log ports.Logger, client *redis.Client) (handle *Handle) {

	handle = &Handle{
		log:    log,
		client: client,
	}

	return
}

// Create all arguments of Redis Set like nx, xx etc.
func (h *Handle) Create(ctx context.Context, i ...interface{}) (*models.Response, error) {

	paramsLength := len(i)
	if paramsLength < 2 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err error

		key        string
		value      interface{}
		expiration time.Duration

		isKey, isExpiration bool
	)

	if key, isKey = i[0].(string); !isKey {
		err = utilities.NewError(utilities.InvalidParameter.Error(), "key")
		return nil, err
	}

	value = i[1]

	if paramsLength > 2 {
		if expiration, isExpiration = i[2].(time.Duration); !isExpiration {
			err = utilities.NewError(utilities.InvalidParameter.Error(), "expiration")
			return nil, err
		}
	}

	resultCmd := h.client.Set(ctx, key, value, expiration)
	if err = resultCmd.Err(); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), err.Error())
	}

	var result string
	if result, err = resultCmd.Result(); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), err.Error())
	}

	return &models.Response{
		Result: []interface{}{result},
	}, nil
}

func (h *Handle) Query(ctx context.Context, i ...interface{}) (*models.Response, error) {

	paramsLength := len(i)
	if paramsLength < 2 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err error

		key   string
		isKey bool
	)

	if key, isKey = i[0].(string); !isKey {
		err = utilities.NewError(utilities.InvalidParameter.Error(), "key")
		return nil, err
	}

	resultCmd := h.client.Get(ctx, key)
	if err = resultCmd.Err(); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), err.Error())
	}

	var result string
	if result, err = resultCmd.Result(); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), err.Error())
	}

	return &models.Response{
		Result: []interface{}{result},
	}, nil
}

func (h *Handle) Update(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return h.Create(ctx, i...)
}

func (h *Handle) Delete(ctx context.Context, i ...interface{}) (*models.Response, error) {

	paramsLength := len(i)
	if paramsLength < 2 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err error

		key   string
		isKey bool
	)

	if key, isKey = i[0].(string); !isKey {
		err = utilities.NewError(utilities.InvalidParameter.Error(), "key")
		return nil, err
	}

	resultCmd := h.client.Del(ctx, key)
	if err = resultCmd.Err(); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), err.Error())
	}

	var result int64
	if result, err = resultCmd.Result(); err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), err.Error())
	}

	return &models.Response{
		Result: []interface{}{result},
	}, nil
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
