package gormx

// Delete deletes the record from database by the given conditions.
func Delete[T any](where map[any]any) (err error) {
	var f T
	err = GetDB().First(&f, where).Error
	if err != nil {
		return
	}

	err = GetDB().Delete(&f).Error
	return
}
