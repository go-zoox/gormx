package gormx

// FindOneOrCreate find one or create one.
func FindOneOrCreate[T any](where map[any]any, callback func(*T)) (*T, error) {
	f, err := FindOne[T](where)
	if err != nil {
		var tmp T
		callback(&tmp)

		if f, err = Create(&tmp); err != nil {
			return nil, err
		}
	}

	return f, nil
}
