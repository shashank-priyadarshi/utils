package sql

import (
	"context"
	"database/sql"

	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Handle struct {
	log    ports.Logger
	client *sql.DB
}

func NewMySQLHandle(log ports.Logger, client *sql.DB) (handle *Handle) {
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
