package gormx

// Delete deletes the record from database by the given conditions.
// Supports both map[any]any and *Where as where condition.
func Delete[T any, W WhereCondition](where W) (err error) {
	// Use FindOne with generic where condition
	f, err := FindOne[T](where)
	if err != nil {
		return err
	}

	err = GetDB().Delete(f).Error
	return
}
