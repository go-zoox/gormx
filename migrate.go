package gormx

import "github.com/go-zoox/logger"

// Migrate migrates the models to the database.
func Migrate() {
	if model == nil {
		panic("models must be register first")
	}

	db := GetDB()

	total := model.Length()
	current := 0
	logger.Infof("[gormx][migrate] models total: %d", total)
	model.ForEach(func(id string, service any) {
		current++
		logger.Infof("[gormx][migrate] migrate: %d/%d ...", current, total)
		db.AutoMigrate(service)
	})
}
