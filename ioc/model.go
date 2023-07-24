package ioc

import (
	"fmt"
	"time"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
	"gorm.io/gorm"
)

type Model interface {
	ModelName() string
	Model() container.Container
}

func RegisterModel(name string, m Model) {
	if model.Has(name) {
		panic("model already exists: " + name)
	}

	logger.Info("[cms][model] register: %s", name)
	model.Register(name, m)
}

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

type ModelImpl struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//
	Creator  uint `json:"creator"`
	Modifier uint `json:"modifier"`
}

func (m *ModelImpl) ModelName() string {
	panic("model.ModelName() not implemented")
}

func (m *ModelImpl) Model() container.Container {
	return model
}
