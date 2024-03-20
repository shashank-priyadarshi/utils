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

func (h *Handler) Create(ctx context.Context, args ...interface{}) (*models.Response, error) {
	return h.Update(ctx, args...)
}

func (h *Handler) Query(ctx context.Context, args ...interface{}) (*models.Response, error) {
	h.client.QueryContext(ctx, "SELECT * FROM something", args...)
	return nil, nil
}

func (h *Handler) Update(ctx context.Context, args ...interface{}) (*models.Response, error) {
	h.client.ExecContext(ctx, "SELECT * FROM TABLE ( something )", args...)
	return nil, nil
}

func (h *Handler) Delete(ctx context.Context, args ...interface{}) (*models.Response, error) {
	return h.Update(ctx, args...)
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
