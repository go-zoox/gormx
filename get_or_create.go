package gormx

// GetOrCreate gets or creates a record.
func GetOrCreate[T any](where map[any]any, callback func(*T)) (*T, error) {
	return FindOneOrCreate(where, callback)
}
