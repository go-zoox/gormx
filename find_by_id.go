package gormx

// FindByID finds a record by id.
func FindByID[T any](id uint) (*T, error) {
	var f T
	if err := GetDB().First(&f, id).Error; err != nil {
		return nil, err
	}

	return &f, nil
}
