package rdbms

import (
	"database/sql"
	"fmt"
	"github.com/shashank-priyadarshi/utilities/database/constants"
	ormHandler "github.com/shashank-priyadarshi/utilities/database/internal/rdbms/orm"
	sqlHandler "github.com/shashank-priyadarshi/utilities/database/internal/rdbms/sql"
	"github.com/shashank-priyadarshi/utilities/database/ports"
	ports2 "github.com/shashank-priyadarshi/utilities/logger/ports"
)

func NewRelationalDBHandle(log ports2.Logger, withORM bool, orm constants.ORM, client interface{}) (handle ports.Database, err error) {

	switch withORM {
	case true:
		handle, err = ormHandler.NewORMHandle(log, orm, client)
		if err != nil {
			err = fmt.Errorf("error creating orm handle: %w", err)
			log.Error(err)
			return
		}

	case false:
		conn, ok := client.(*sql.DB)
		if !ok {
			err = fmt.Errorf("invalid sql connection: %w", err)
			log.Error(err)
			return
		}

		handle = sqlHandler.NewMySQLHandle(log, conn)
	}

	return
}
