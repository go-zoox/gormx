package gormx

import (
	"errors"

	"gorm.io/gorm"
)

// FindOneOrCreate find one or create one.
// Supports both map[any]any and *Where as where condition.
func FindOneOrCreate[T any, W WhereCondition](where W, callback func(*T)) (*T, error) {
	f, err := FindOne[T](where)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		var tmp T
		callback(&tmp)

		if f, err = Create(&tmp); err != nil {
			return nil, err
		}
	}

	return f, nil
}
