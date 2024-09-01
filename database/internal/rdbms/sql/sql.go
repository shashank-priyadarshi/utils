package sql

import (
	"context"
	"database/sql"
	"fmt"
	"go.ssnk.in/utils/errors"

	"go.ssnk.in/utils/database/models"
	"go.ssnk.in/utils/errors"
)

type Handler struct {
	client *sql.DB
}

func Handle(client *sql.DB) *Handler {
	return &Handler{
		client: client,
	}
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
		return nil, errors.InsufficientParameters.Error(2, len(args))
	}

	var (
		err   error
		query string
		ok    bool
		rows  *sql.Rows
		model interface{}
	)

	if model = args[0]; model == nil {
		return nil, errors.InvalidParameterType.Error("model", model, args[0])
	}

	if query, ok = args[1].(string); !ok {
		return nil, errors.InvalidParameterType.Error("query", query, args[1])
	}

	rows, err = h.client.QueryContext(ctx, query, args[1:]...)
	if err != nil {
		return nil, errors.OperationFailed.Error(fmt.Sprintf("error executing query %s on database: %v", query, err))
	}

	var response *models.Response

	for rows.Next() {
		err = rows.Scan(model)
		if err != nil {
			return nil, errors.OperationFailed.Error(fmt.Sprintf("error scanning db row: %v", err))
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
		return nil, errors.InsufficientParameters.Error(1, len(args))
	}

	var (
		err    error
		query  string
		ok     bool
		result sql.Result
	)

	if query, ok = args[0].(string); !ok {
		return nil, errors.InvalidParameterType.Error("query", query, args[0])
	}

	if len(query) == 0 {
		return nil, errors.InvalidParameterValue.Error("query", "valid string", query)
	}

	if len(args) > 1 {
		args = args[1:]
	}

	result, err = h.client.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, errors.OperationFailed.Error(fmt.Sprintf("error executing query %s on database: %v", query, err))
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
