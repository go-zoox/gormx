package ddl

import (
	"fmt"

	"github.com/go-zoox/gormx"
)

func execute(engine, dsn, sql string) error {
	db, err := gormx.Connect(engine, dsn)
	if err != nil {
		return fmt.Errorf("connecting database failed: %s", err.Error())
	}
	conn, err := db.DB()
	if err != nil {
		return fmt.Errorf("getting database connection failed: %s", err.Error())
	}
	defer conn.Close()

	var f interface{}
	return db.Raw(sql).Scan(&f).Error
}
