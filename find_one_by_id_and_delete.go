package gormx

// FindOneByIDAndDelete finds one record by id and delete it.
func FindOneByIDAndDelete[T any](id uint) (*T, error) {
	f, err := FindByID[T](id)
	if err != nil {
		return nil, err
	}

	err = GetDB().Delete(f).Error
	return f, err
}
