package connections

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/glebarez/sqlite"
	"github.com/shashank-priyadarshi/utilities/database/constants"
	"github.com/shashank-priyadarshi/utilities/database/models"
	"github.com/shashank-priyadarshi/utilities/logger/ports"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RDBMS(ctx context.Context, log ports.Logger, config *models.Config) (client interface{}, err error) {

	if !isSupported(config.Options.Driver) {
		err = fmt.Errorf("unsupported sql driver")
		log.Error(err)
		return
	}

	conn, err := sql.Open(string(config.Options.Driver), "")
	if err != nil {
		err = fmt.Errorf("error connecting to database: %v", err)
		log.Error(err)
		return
	}

	client = conn

	switch config.Options.WithORM {
	case true:
		switch config.Options.ORM {
		case constants.GORM:
			var gormDB *gorm.DB
			if gormDB, err = createGORMConnection(config.Options.Driver, conn); err != nil {
				err = fmt.Errorf("error creating gorm connection: %v", err)
				return
			} else {
				client = gormDB
			}
			return

		default:
			err = fmt.Errorf("unsupported orm type")
			log.Error(err)
			return
		}

	default:
		log.Info("no orm configured")
		return
	}
}

func isSupported(driver constants.Driver) bool {

	var supported = make(map[constants.Driver]any)

	supported[constants.VITESS] = nil
	supported[constants.COCKROACHDB] = nil

	// GORM supported
	supported[constants.MYSQL] = nil
	supported[constants.POSTGRES] = nil
	supported[constants.SQLITE] = nil

	if _, ok := supported[driver]; !ok {
		return false
	}

	return true
}

func createGORMConnection(driver constants.Driver, conn *sql.DB) (db *gorm.DB, err error) {

	switch driver {
	case constants.MYSQL:
		db, err = gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{})
		return

	case constants.POSTGRES:
		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: conn,
		}), &gorm.Config{})
		return

	case constants.SQLITE:
		db, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
		return

	default:
		err = fmt.Errorf("unsupported gorm connection for the driver type %s", driver)
	}

	return
}
