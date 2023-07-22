package gormx

// FindOneByIDOrCreate finds one record by id or create a new one.
func FindOneByIDOrCreate[T any](id uint, callback func(*T)) (*T, error) {
	f, err := FindByID[T](id)
	if err != nil {
		var tmp T
		callback(&tmp)

		if f, err = Create(&tmp); err != nil {
			return nil, err
		}
	}

	return f, nil
}
