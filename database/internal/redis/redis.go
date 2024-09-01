package redis

import (
	"context"
	"go.ssnk.in/utils/errors"
	"time"

	"github.com/redis/go-redis/v9"
	"go.ssnk.in/utils/database/models"
	"go.ssnk.in/utils/errors"
)

type Handler struct {
	client *redis.Client
}

func Handle(client *redis.Client) (handle *Handler) {
	handle = &Handler{
		client: client,
	}

	return
}

// Create all arguments of Redis Set like nx, xx etc.
func (h *Handler) Create(ctx context.Context, args ...interface{}) (*models.Response, error) {
	paramsLength := len(args)
	if paramsLength < 2 {
		return nil, errors.InsufficientParameters.Error(2, paramsLength)
	}

	var (
		err error

		key        string
		value      interface{}
		expiration time.Duration

		isKey, isExpiration bool
	)

	if key, isKey = args[0].(string); !isKey {
		err = errors.InvalidParameterType.Error("key", key, args[0])
		return nil, err
	}

	value = args[1]

	if paramsLength > 2 {
		if expiration, isExpiration = args[2].(time.Duration); !isExpiration {
			err = errors.InvalidParameterType.Error("expiration", expiration, args[2])
			return nil, err
		}
	}

	resultCmd := h.client.Set(ctx, key, value, expiration)
	if err = resultCmd.Err(); err != nil {
		return nil, errors.OperationFailed.Error(err.Error())
	}

	var result string
	if result, err = resultCmd.Result(); err != nil {
		return nil, errors.OperationFailed.Error(err.Error())
	}

	return &models.Response{
		Result: []interface{}{result},
	}, nil
}

func (h *Handler) Query(ctx context.Context, args ...interface{}) (*models.Response, error) {
	paramsLength := len(args)
	if paramsLength < 1 {
		return nil, errors.InsufficientParameters.Error(1, paramsLength)
	}

	var (
		err error

		key   string
		isKey bool
	)

	if key, isKey = args[0].(string); !isKey {
		err = errors.InvalidParameterType.Error("key", key, args[0])
		return nil, err
	}

	resultCmd := h.client.Get(ctx, key)
	if err = resultCmd.Err(); err != nil {
		return nil, errors.OperationFailed.Error(err.Error())
	}

	var result string
	if result, err = resultCmd.Result(); err != nil {
		return nil, errors.OperationFailed.Error(err.Error())
	}

	return &models.Response{
		Result: []interface{}{result},
	}, nil
}

func (h *Handler) Update(ctx context.Context, args ...interface{}) (*models.Response, error) {
	return h.Create(ctx, args...)
}

func (h *Handler) Delete(ctx context.Context, args ...interface{}) (*models.Response, error) {
	paramsLength := len(args)
	if paramsLength < 1 {
		return nil, errors.InsufficientParameters.Error(1, paramsLength)
	}

	var (
		err error

		key   string
		isKey bool
	)

	if key, isKey = args[0].(string); !isKey {
		err = errors.InvalidParameterType.Error("key", key, args[0])
		return nil, err
	}

	resultCmd := h.client.Del(ctx, key)
	if err = resultCmd.Err(); err != nil {
		return nil, errors.OperationFailed.Error(err.Error())
	}

	var result int64
	if result, err = resultCmd.Result(); err != nil {
		return nil, errors.OperationFailed.Error(err.Error())
	}

	return &models.Response{
		Result: []interface{}{result},
	}, nil
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
