package gormx

// SQL finds one record by id or create a new one.
func SQL[T any](sql string, values ...any) (*T, error) {
	var f T
	if err := GetDB().Raw(sql, values...).Scan(&f).Error; err != nil {
		return nil, err
	}

	return &f, nil
}
