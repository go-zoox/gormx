package ddl

import (
	"fmt"

	"github.com/go-zoox/gormx"
)

// CreateUser creates a user.
func CreateUser(engine, dsn, database, username, password string) (err error) {
	switch engine {
	case "postgres":
		err = execute(engine, dsn, fmt.Sprintf(`CREATE USER %s WITH ENCRYPTED PASSWORD '%s'`, username, password))
		if err != nil {
			return
		}

		err = execute(engine, dsn, fmt.Sprintf(`GRANT ALL PRIVILEGES ON DATABASE %s TO %s`, database, username))
		if err != nil {
			return
		}
	case "mysql":
		// _, err = gormx.SQL[any](`CREATE USER ? @'%' IDENTIFIED BY ?`, username, password)
		err = execute(engine, dsn, fmt.Sprintf(`CREATE USER '%s'@'%%' IDENTIFIED BY '%s'`, username, password))
		if err != nil {
			return
		}

		err = execute(engine, dsn, fmt.Sprintf(`GRANT ALL PRIVILEGES ON %s.* TO '%s'@'%%';`, database, username))
		if err != nil {
			return
		}

		err = execute(engine, dsn, `flush privileges;`)
		if err != nil {
			return
		}
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}

// DeleteUser deletes a user.
func DeleteUser(engine, dsn, username string) (err error) {
	err = execute(engine, dsn, fmt.Sprintf(`DROP USER '%s'`, username))
	return
}

// UpdateUserPassword updates a user's password.
func UpdateUserPassword(username, password string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`ALTER USER ? WITH ENCRYPTED PASSWORD ?`, username, password)
	case "mysql":
		_, err = gormx.SQL[any](`ALTER USER ? @'%' IDENTIFIED BY ?`, username, password)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}

// GrantUserPrivileges grants privileges to a user.
func GrantUserPrivileges(username string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO ?`, username)
	case "mysql":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON *.* TO ? @'%' WITH GRANT OPTION`, username)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}

// RevokeUserPrivileges revokes privileges from a user.
func RevokeUserPrivileges(username string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`REVOKE ALL PRIVILEGES ON ALL TABLES IN SCHEMA public FROM ?`, username)
	case "mysql":
		_, err = gormx.SQL[any](`REVOKE ALL PRIVILEGES ON *.* FROM ? @'%'`, username)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}

// GrantUserPrivilegesToDatabase grants privileges to a user on a database.
func GrantUserPrivilegesToDatabase(username, database string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON DATABASE ? TO ?`, database, username)
	case "mysql":
		_, err = gormx.SQL[any](`GRANT ALL PRIVILEGES ON ?.* TO ? @'%'`, database, username)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}

// RevokeUserPrivilegesFromDatabase revokes privileges from a user on a database.
func RevokeUserPrivilegesFromDatabase(username, database string) (err error) {
	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`REVOKE ALL PRIVILEGES ON DATABASE ? FROM ?`, database, username)
	case "mysql":
		_, err = gormx.SQL[any](`REVOKE ALL PRIVILEGES ON ?.* FROM ? @'%'`, database, username)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}

// ReadOnlyUserToDatabase grants read-only privileges to a user on a database.
func ReadOnlyUserToDatabase(username, database string) (err error) {
	if err = GrantUserPrivilegesToDatabase(username, database); err != nil {
		return
	}

	switch gormx.GetEngine() {
	case "postgres":
		_, err = gormx.SQL[any](`GRANT SELECT ON ALL TABLES IN SCHEMA public TO ?`, username)
	case "mysql":
		_, err = gormx.SQL[any](`GRANT SELECT ON ?.* TO ? @'%'`, database, username)
	default:
		err = fmt.Errorf("unsupported engine: %s", gormx.GetEngine())
	}

	return
}
