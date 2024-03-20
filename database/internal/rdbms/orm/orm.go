package orm

import (
	"fmt"

	// go run -mod=mod entgo.io/ent/cmd/ent generate --template <dir-path> --template glob="path/to/*.tmpl" ./ent/schema
	_ "entgo.io/ent/entc"
	_ "entgo.io/ent/entc/gen"
	_ "entgo.io/ent/schema/field"
	"github.com/shashank-priyadarshi/utilities/database/constants"
	gormHandler "github.com/shashank-priyadarshi/utilities/database/internal/rdbms/orm/gorm"
	"github.com/shashank-priyadarshi/utilities/database/ports"
	ports2 "github.com/shashank-priyadarshi/utilities/logger/ports"
	"gorm.io/gorm"
)

func NewORMHandle(log ports2.Logger, orm constants.ORM, client interface{}) (handle ports.Database, err error) {

	switch orm {
	case constants.GORM:
		conn, ok := client.(*gorm.DB)
		if !ok {
			err = fmt.Errorf("invalid gorm connection: %w", err)
			log.Error(err)
			return
		}

		handle = gormHandler.NewGORMHandle(log, conn)
		return
	}

	return
}
