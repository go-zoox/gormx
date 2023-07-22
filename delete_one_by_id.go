package gormx

// DeleteOneByID deletes one record by id.
func DeleteOneByID[T any](id uint) (err error) {
	var f T
	err = GetDB().First(&f, id).Error
	if err != nil {
		return
	}

	err = GetDB().Delete(&f).Error
	return
}
