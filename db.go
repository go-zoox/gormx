package gormx

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

var metadataEngine string
var metadataDSN string

// LoadDBOptions is the options for LoadDB
type LoadDBOptions struct {
	IsProd bool
	//
	TablePrefix string
	//
	DryRun bool
}

// GetDB returns the gorm.DB instance
func GetDB() *gorm.DB {
	if db == nil {
		panic("DB is not initialized")
	}

	return db
}

// GetEngine returns the database engine
func GetEngine() string {
	return metadataEngine
}

// GetDSN returns the database DSN
func GetDSN() string {
	return metadataDSN
}

// LoadDB loads the database
func LoadDB(engine string, dsn string, opts ...func(*LoadDBOptions)) (err error) {
	opt := &LoadDBOptions{}
	for _, o := range opts {
		o(opt)
	}

	var dialector gorm.Dialector
	switch engine {
	case "postgres":
		dialector = postgres.Open(dsn)
	case "mysql":
		dialector = mysql.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		return fmt.Errorf("unknown engine: %s", engine)
	}

	metadataEngine = engine
	metadataDSN = dsn

	logLevel := logger.Info
	if opt.IsProd {
		logLevel = logger.Error
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   opt.TablePrefix,
		},
		Logger:               logger.Default.LogMode(logLevel), // Print SQL queries
		DisableAutomaticPing: false,
		// DisableForeignKeyConstraintWhenMigrating: true,
		DryRun: opt.DryRun,
	})
	if err != nil {
		return fmt.Errorf("connecting database failed: %s", err.Error())
	}

	return nil
}
