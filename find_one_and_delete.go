package gormx

// FindOneAndDelete finds one record and delete it.
// Supports both map[any]any and *Where as where condition.
func FindOneAndDelete[T any, W WhereCondition](where W) (*T, error) {
	f, err := FindOne[T](where)
	if err != nil {
		return nil, err
	}

	err = GetDB().Delete(f).Error
	return f, err
}
