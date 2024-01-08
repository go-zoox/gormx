package gormx

import (
	"fmt"
	"sync"

	"github.com/go-zoox/ioc"
	"github.com/go-zoox/logger"
)

var model ioc.Container
var once = &sync.Once{}

// Model is the interface that wraps the basic methods.
type Model interface {
	ModelName() string
	Model() ioc.Container
}

// Register registers the model.
func Register(name string, m Model) {
	once.Do(func() {
		model = ioc.New()
	})

	if model.Has(name) {
		panic("model already exists: " + name)
	}

	logger.Infof("[cms][model] register: %s", name)
	model.Register(name, m)
}

// Get returns the model by the given id.
func Get[T any](id string) T {
	if !model.Has(id) {
		panic("model not registered: " + id)
	}

	s, ok := model.MustGet(id).(T)
	if !ok {
		panic(fmt.Sprintf("model not valid type(%v): %s", new(T), id))
	}

	return s
}
