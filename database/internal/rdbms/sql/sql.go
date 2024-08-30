package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/shashank-priyadarshi/utilities"

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

/*
Query

Add query validator to validate query and required arguments
*/
func (h *Handler) Query(ctx context.Context, args ...interface{}) (*models.Response, error) {

	if len(args) < 2 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err   error
		query string
		ok    bool
		rows  *sql.Rows
		model interface{}
	)

	if model = args[0]; model == nil {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "model")
	}

	if query, ok = args[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "query")
	}

	rows, err = h.client.QueryContext(ctx, query, args[1:]...)
	if err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error executing query %s on database: %v", query, err))
	}

	var response *models.Response

	for rows.Next() {
		err = rows.Scan(model)
		if err != nil {
			return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error scanning db row: %v", err))
		}

		response.Result = append(response.Result, model)
	}

	return response, nil
}

/*
Update

Add query validator to validate query and required arguments
*/
func (h *Handler) Update(ctx context.Context, args ...interface{}) (*models.Response, error) {

	if len(args) < 1 {
		return nil, utilities.InsufficientParameters
	}

	var (
		err    error
		query  string
		ok     bool
		result sql.Result
	)

	if query, ok = args[1].(string); !ok {
		return nil, utilities.NewError(utilities.InvalidParameter.Error(), "query")
	}

	result, err = h.client.ExecContext(ctx, query, args[1:]...)
	if err != nil {
		return nil, utilities.NewError(utilities.OperationFailed.Error(), fmt.Sprintf("error executing query %s on database: %v", query, err))
	}

	var response *models.Response

	response.Result = []interface{}{
		&result,
	}

	return response, nil
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
