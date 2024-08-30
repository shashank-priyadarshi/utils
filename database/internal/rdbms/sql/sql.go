package sql

import (
	"context"
	"database/sql"

	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Handler struct {
	log    ports.Logger
	client *sql.DB
}

func Handle(log ports.Logger, client *sql.DB) (handle *Handler) {
	handle.log = log
	handle.client = client

	return
}

func (h *Handler) Create(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Query(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Update(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Delete(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Begin(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Execute(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Rollback(ctx context.Context, i ...interface{}) (*models.Response, error) {
	return nil, nil
}

func (h *Handler) Configure(ctx context.Context, i ...interface{}) error {
	return nil
}

func (h *Handler) Close() error {
	return nil
}
