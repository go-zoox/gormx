package gormx

// Save saves a record.
func Save[T any](one *T) error {
	return GetDB().Save(one).Error
}
