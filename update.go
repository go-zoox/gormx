package gormx

// Update updates a record.
func Update[T any](id uint, uc func(*T)) (err error) {
	var f T
	err = GetDB().First(&f, id).Error
	if err != nil {
		return
	}

	uc(&f)

	err = GetDB().Save(&f).Error
	return
}
