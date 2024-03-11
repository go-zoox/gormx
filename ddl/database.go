package ddl

import (
	"context"
	"fmt"

	"github.com/go-zoox/gormx"
	"github.com/go-zoox/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateDatabase creates a database.
func CreateDatabase(engine, dsn, name string) (err error) {
	switch engine {
	case "postgres":
		err = execute(engine, dsn, fmt.Sprintf("CREATE DATABASE %s", name))
	case "mysql":
		err = execute(engine, dsn, fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARSET 'utf8'", name))
	case "mongodb":
		logger.Infof("MongoDB does not need to create a database")
	default:
		err = fmt.Errorf("unsupported engine: %s, available engines: postgres, mysql", engine)
	}
	return
}

// DeleteDatabase deletes a database.
func DeleteDatabase(engine, dsn, name string) (err error) {
	switch engine {
	case "postgres":
		err = execute(engine, dsn, fmt.Sprintf("DROP DATABASE %s", name))
	case "mysql":
		err = execute(engine, dsn, fmt.Sprintf("DROP DATABASE %s", name))
	case "mongodb":
		err = executeMongoDB(dsn, func(ctx context.Context, client *mongo.Client) error {
			return client.Database(name).Drop(ctx)
		})
	default:
		err = fmt.Errorf("unsupported engine: %s, available engines: postgres, mysql", engine)
	}
	return
}

// AddUserToDatabase adds a user to a database.
func AddUserToDatabase(engine, username, database string) (err error) {
	switch engine {
	case "postgres":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON ? TO ?`, database, username)
	case "mysql":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON ?.* TO ? @'%'; FLUSH PRIVILEGES;`, database, username)
	default:
		err = fmt.Errorf("unsupported engine: %s, available engines: postgres, mysql", engine)
	}

	return
}

// RemoveUserFromDatabase removes a user from a database.
func RemoveUserFromDatabase(username, database string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`REVOKE ALL PRIVILEGES ON ? FROM ?`, database, username)
	case "mysql":
		_, err = gormx.SQL[any](`REVOKE ALL PRIVILEGES ON ?.* FROM ? @'%'; FLUSH PRIVILEGES;`, database, username)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}
