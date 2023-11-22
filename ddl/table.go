package ddl

import "github.com/go-zoox/gormx"

// CreateTableInDatabase creates a table in a database.
func CreateTableInDatabase(database, name string) (err error) {
	_, err = gormx.SQL[any](`USE ?; CREATE TABLE ?`, database, name)
	return
}

// DeleteTableInDatabase deletes a table in a database.
func DeleteTableInDatabase(database, name string) (err error) {
	_, err = gormx.SQL[any](`USE ?; DROP TABLE ?`, database, name)
	return
}

// RenameTableInDatabase renames a table in a database.
func RenameTableInDatabase(database, oldName, newName string) (err error) {
	_, err = gormx.SQL[any](`USE ?; RENAME TABLE ? TO ?`, database, oldName, newName)
	return
}
