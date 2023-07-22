package gormx

// FindOneByIDAndUpdate finds one by id and update.
func FindOneByIDAndUpdate[T any](id uint, callback func(*T)) (*T, error) {
	f, err := FindByID[T](id)
	if err != nil {
		return nil, err
	}

	callback(f)

	return f, nil
}
