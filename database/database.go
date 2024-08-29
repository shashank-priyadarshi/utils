package database

import (
	"context"
	"fmt"
	"go.ssnk.in/utils/database/constants"
	adapters "go.ssnk.in/utils/database/internal"
	"go.ssnk.in/utils/database/models"
	"go.ssnk.in/utils/database/ports"
)

type Database struct {
	ports.Database
}

func New(ctx context.Context, config *models.Config) (Database, error) {

	if !isSupported(config.Type) {
		return Database{}, fmt.Errorf("unsupported database type: %s", config.Type)
	}

	db, err := adapters.New(ctx, config)

	return Database{db}, err
}

func isSupported(db constants.Database) bool {

	var supported = make(map[constants.Database]any)
	supported[constants.MYSQLDB] = nil
	supported[constants.MONGODB] = nil
	supported[constants.REDIS] = nil

	if _, ok := supported[db]; !ok {
		return false
	}

	return true
}
