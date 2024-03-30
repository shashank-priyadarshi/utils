package rdbms

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/log"

	"github.com/shashank-priyadarshi/utilities/database/constants"
	ormhandler "github.com/shashank-priyadarshi/utilities/database/internal/rdbms/orm"
	sqlhandler "github.com/shashank-priyadarshi/utilities/database/internal/rdbms/sql"
	"github.com/shashank-priyadarshi/utilities/database/ports"
)

func Handle(withORM bool, orm constants.ORM, client interface{}) (handle ports.Database, err error) {

	switch withORM {
	case true:
		handle, err = ormhandler.Handle(orm, client)
		if err != nil {
			err = fmt.Errorf("error creating orm handle: %w", err)
			log.Error(err)
			return
		}

	case false:
		conn, ok := client.(*sql.DB)
		if !ok {
			err = fmt.Errorf("invalid sql connection: %w", err)
			return
		}

		handle = sqlhandler.Handle(conn)
	}

	return
}
