package gormx

import (
	"errors"

	"gorm.io/gorm"
)

// Exists returns true if the record exists.
// Supports both map[any]any and *Where as where condition.
func Exists[T any, W WhereCondition](where W) (bool, error) {
	_, err := FindOne[T](where)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
