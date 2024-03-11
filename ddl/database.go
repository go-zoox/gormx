package ddl

import (
	"fmt"

	"github.com/go-zoox/gormx"
)

// CreateDatabase creates a database.
func CreateDatabase(engine, dsn, name string) (err error) {
	switch engine {
	case "postgres":
		err = execute(engine, dsn, fmt.Sprintf("CREATE DATABASE %s", name))
	case "mysql":
		err = execute(engine, dsn, fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARSET 'utf8'", name))
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
