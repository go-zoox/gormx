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

		return db.AutoMigrate(s)
	})
	if err != nil {
		panic(fmt.Errorf("failed to migrate: %s", err))
	}
}
