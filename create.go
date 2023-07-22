package gormx

// Create creates a record.
func Create[T any](one *T) (*T, error) {
	err := GetDB().
		Create(one).Error

	return one, err
}
