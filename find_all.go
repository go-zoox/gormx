package gormx

// FindAll finds all records.
func FindAll[T any](where *Where, orderBy *OrderBy) (data []*T, err error) {
	return ListALL[T](where, orderBy)
}
