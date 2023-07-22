package gormx

// Retrieve retrieves a record.
func Retrieve[T any](id uint) (*T, error) {
	var f T
	if err := GetDB().First(&f, id).Error; err != nil {
		return nil, err
	}

	return &f, nil
}
