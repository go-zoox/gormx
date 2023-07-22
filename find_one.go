package gormx

// FindOne finds one record.
func FindOne[T any](where map[any]any) (*T, error) {
	var f T
	if err := GetDB().First(&f, where).Error; err != nil {
		return nil, err
	}

	return &f, nil
}
