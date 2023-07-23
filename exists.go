package gormx

import (
	"errors"

	"gorm.io/gorm"
)

// Exists returns true if the record exists.
func Exists[T any](where map[any]any) (bool, error) {
	_, err := FindOne[T](where)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
