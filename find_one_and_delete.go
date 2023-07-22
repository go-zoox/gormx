package gormx

// FindOneAndDelete finds one record and delete it.
func FindOneAndDelete[T any](where map[any]any) (*T, error) {
	f, err := FindOne[T](where)
	if err != nil {
		return nil, err
	}

	err = GetDB().Delete(f).Error
	return f, err
}
