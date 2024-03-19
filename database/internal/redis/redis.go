package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Handle struct {
	log    ports.Logger
	client *redis.Client
}

func NewRedisHandle(log ports.Logger, client *redis.Client) (handle *Handle) {
	handle.log = log
	handle.client = client

	return
}

func (h *Handle) Create(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handle) Query(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handle) Update(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
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
