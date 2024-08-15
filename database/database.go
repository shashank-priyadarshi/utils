package database

import (
	"context"
	"fmt"
	"github.com/shashank-priyadarshi/utilities/database/constants"
	adapters "github.com/shashank-priyadarshi/utilities/database/internal"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/database/ports"
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
