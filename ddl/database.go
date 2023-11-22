package ddl

import (
	"fmt"

	"github.com/go-zoox/gormx"
)

// CreateDatabase creates a database.
func CreateDatabase(name string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`CREATE DATABASE ?`, name)
	case "mysql":
		_, err = gormx.SQL[any](`CREATE DATABASE ? DEFAULT CHARSET 'utf8'`, name)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}
	return
}

// DeleteDatabase deletes a database.
func DeleteDatabase(name string) (err error) {
	_, err = gormx.SQL[any](`DROP DATABASE ?`, name)
	return
}

// AddUserToDatabase adds a user to a database.
func AddUserToDatabase(username, database string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON ? TO ?`, database, username)
	case "mysql":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON ?.* TO ? @'%'; FLUSH PRIVILEGES;`, database, username)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
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
