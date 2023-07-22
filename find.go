package gormx

// Find finds records.
func Find[T any](page, pageSize uint, where *Where, orderBy *OrderBy) (data []*T, total int64, err error) {
	return List[T](page, pageSize, where, orderBy)
}
