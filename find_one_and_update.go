package gormx

// FindOneAndUpdate finds one and update it.
// Supports both map[any]any and *Where as where condition.
func FindOneAndUpdate[T any, W WhereCondition](where W, callback func(*T)) (*T, error) {
	f, err := FindOne[T](where)
	if err != nil {
		return nil, err
	}

	callback(f)

	return f, nil
}
