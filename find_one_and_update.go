package gormx

// FindOneAndUpdate finds one and update it.
func FindOneAndUpdate[T any](where map[any]any, callback func(*T)) (*T, error) {
	var f T
	if err := GetDB().First(&f, where).Error; err != nil {
		return nil, err
	}

	callback(&f)

	return &f, nil
}
