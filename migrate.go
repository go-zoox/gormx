package gormx

import (
	"fmt"

	"github.com/go-zoox/logger"
)

// Migrate migrates the models to the database.
func Migrate() {
	if model == nil {
		panic("models must be register first")
	}

	// db := GetDB()

	total := model.Length()
	current := 0
	logger.Infof("[gormx][migrate] models total: %d", total)
	err := model.ForEach(func(id string, s any) error {
		current++
		logger.Infof("[gormx][migrate][%d/%d] migrate: %s ...", current, total, s.(Model).ModelName())

		// @TODO
		// bug:  ERROR: constraint "uni_v1_devops_dict_uuid" of relation "v1_devops_dict" does not exist (SQLSTATE 42704)
		// issue: https://github.com/go-gorm/gorm/issues/7010
		//
		// fix:
		//   ALTER TABLE v1_devops_dict DROP CONSTRAINT idx_v1_devops_dict_uuid;
		//
		return db.AutoMigrate(s)

		// db.AutoMigrate(s)
		// return nil
	})
	if err != nil {
		panic(fmt.Errorf("failed to migrate: %s", err))
	}
}
