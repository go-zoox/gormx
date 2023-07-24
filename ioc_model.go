package gormx

import (
	"fmt"
	"time"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
	"gorm.io/gorm"
)

// Model is the interface that wraps the basic methods.
type Model interface {
	ModelName() string
	Model() container.Container
}

// RegisterModel registers a model.
func RegisterModel(name string, m Model) {
	if model.Has(name) {
		panic("model already exists: " + name)
	}

	logger.Info("[cms][model] register: %s", name)
	model.Register(name, m)
}

// GetModel returns the model by the given id.
func GetModel[T any](id string) T {
	if !model.Has(id) {
		panic("model not registered: " + id)
	}

	s, ok := model.MustGet(id).(T)
	if !ok {
		panic(fmt.Sprintf("model not valid type(%v): %s", new(T), id))
	}

	return s
}

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
func (m *ModelImpl) Model() container.Container {
	return model
}
