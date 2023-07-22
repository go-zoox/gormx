package gormx

// GetMany gets many records by ids.
func GetMany[T any](ids []uint) (data []*T, err error) {
	err = GetDB().
		Where("id IN (?)", ids).
		Find(&data).Error
	return
}
