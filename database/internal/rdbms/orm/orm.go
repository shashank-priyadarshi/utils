package orm

import (
	"fmt"

	// go run -mod=mod entgo.io/ent/cmd/ent generate --template <dir-path> --template glob="path/to/*.tmpl" ./ent/schema
	_ "entgo.io/ent/entc"
	_ "entgo.io/ent/entc/gen"
	_ "entgo.io/ent/schema/field"
	"go.ssnk.in/utils/database/constants"
	gormhandler "go.ssnk.in/utils/database/internal/rdbms/orm/gorm"
	"go.ssnk.in/utils/database/ports"
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
