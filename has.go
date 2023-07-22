package gormx

// Has returns true if the record exists.
func Has[T any](where map[string]any) bool {
	var f T
	if err := GetDB().First(&f, where).Error; err != nil {
		return false
	}

	return true
}
