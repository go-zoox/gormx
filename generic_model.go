package gormx

import (
	"time"

	"github.com/go-zoox/ioc"
	"gorm.io/gorm"
)

// ModelImpl is the implementation of the Model interface.
type ModelImpl struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//
	Creator  uint `json:"creator"`
	Modifier uint `json:"modifier"`
}

// ModelName returns the name of the model.
func (m *ModelImpl) ModelName() string {
	panic("model.ModelName() not implemented")
}

// Model returns the model container.
func (m *ModelImpl) Model() ioc.Container {
	return model
}

// ModelGeneric ...
type ModelGeneric[T any] struct {
}

// List ...
func (m *ModelGeneric[T]) List(page, pageSize uint, where *Where, orderBy *OrderBy) (data []*T, total int64, err error) {
	return List[T](page, pageSize, where, orderBy)
}

// Create ...
func (m *ModelGeneric[T]) Create(one *T) (*T, error) {
	return Create(one)
}

// Retrieve ...
func (m *ModelGeneric[T]) Retrieve(id uint) (*T, error) {
	return Retrieve[T](id)
}

// Update ...
func (m *ModelGeneric[T]) Update(id uint, uc func(*T)) (err error) {
	return Update(id, uc)
}

// Delete ...
func (m *ModelGeneric[T]) Delete(id uint) (err error) {
	return DeleteOneByID[T](id)
}

// Save ...
func (m *ModelGeneric[T]) Save() error {
	return Save(m)
}

// GetMany ...
func (m *ModelGeneric[T]) GetMany(ids []uint) (data []*T, err error) {
	return GetMany[T](ids)
}

// Exists ...
func (m *ModelGeneric[T]) Exists(where map[any]any) (bool, error) {
	return Exists[*T](where)
}

// FindByID ...
func (m *ModelGeneric[T]) FindByID(id uint) (*T, error) {
	return FindByID[T](id)
}

// FindOne ...
func (m *ModelGeneric[T]) FindOne(where map[any]any) (*T, error) {
	return FindOne[T](where)
}

// FindAll ...
func (m *ModelGeneric[T]) FindAll(where *Where, orderBy *OrderBy) ([]*T, error) {
	return FindAll[T](where, orderBy)
}

// FindOneOrCreate ...
func (m *ModelGeneric[T]) FindOneOrCreate(where map[any]any, callback func(*T)) (*T, error) {
	return FindOneOrCreate[T](where, callback)
}

// FindOneAndUpdate ...
func (m *ModelGeneric[T]) FindOneAndUpdate(where map[any]any, callback func(*T)) (*T, error) {
	return FindOneAndUpdate(where, callback)
}

// FindOneAndDelete ...
func (m *ModelGeneric[T]) FindOneAndDelete(where map[any]any) (*T, error) {
	return FindOneAndDelete[T](where)
}
