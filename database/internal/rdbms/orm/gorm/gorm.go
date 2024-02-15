package gorm

import (
	"context"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
	"gorm.io/gorm"
)

type Handle struct {
	log    ports.Logger
	client *gorm.DB
}

func NewGORMHandle(log ports.Logger, client *gorm.DB) (handle *Handle) {
	handle.log = log
	handle.client = client

	return
}

func (h *Handle) Create(ctx context.Context, i ...interface{}) (models.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Query(ctx context.Context, i ...interface{}) (models.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Update(ctx context.Context, i ...interface{}) (models.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Delete(ctx context.Context, i ...interface{}) (models.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Begin(ctx context.Context, i ...interface{}) (models.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Execute(ctx context.Context, i ...interface{}) (models.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Rollback(ctx context.Context, i ...interface{}) (models.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Configure(ctx context.Context, i ...interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handle) Close() error {
	//TODO implement me
	panic("implement me")
}
