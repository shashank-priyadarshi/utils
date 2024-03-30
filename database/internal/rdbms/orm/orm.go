package orm

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/database/constants"
	gormhandler "github.com/shashank-priyadarshi/utilities/database/internal/rdbms/orm/gorm"
	"github.com/shashank-priyadarshi/utilities/database/ports"
	"gorm.io/gorm"
)

func Handle(orm constants.ORM, client interface{}) (handle ports.Database, err error) {

	switch orm {
	case constants.GORM:
		conn, ok := client.(*gorm.DB)
		if !ok {
			err = fmt.Errorf("invalid gorm connection: %w", err)
			return
		}

		handle = gormhandler.Handle(conn)
		return
	}

	return
}
