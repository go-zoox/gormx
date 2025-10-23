package gormx

// GetOrCreate gets or creates a record.
// Supports both map[any]any and *Where as where condition.
func GetOrCreate[T any, W WhereCondition](where W, callback func(*T)) (*T, error) {
	return FindOneOrCreate[T](where, callback)
}
