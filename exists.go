package gormx

// Exists returns true if the record exists.
func Exists[T any](where map[any]any) (bool, error) {
	_, err := FindOne[T](where)
	if err != nil {
		return false, err
	}

	return true, nil
}
