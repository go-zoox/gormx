package gormx

// Count counts records.
func Count[T any](where *Where) (count int64, err error) {
	whereClause, whereValues := where.Build()
	countTx := GetDB().Model(new(T))

	if whereClause != "" {
		countTx = countTx.Where(whereClause, whereValues...)
	}

	err = countTx.
		Count(&count).
		Error
	return
}

// CountALL counts all records.
func CountALL[T any]() (total int64, err error) {
	err = GetDB().Model(new(T)).
		Count(&total).
		Error
	return
}
