package gormx

import (
	"fmt"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
)

// Controller is the interface that wraps the basic methods.
type Controller interface {
	Name() string
	//
	Service() container.Container
	//
}

// ControllerImpl is the implementation of the Controller interface.
type ControllerImpl struct {
}

// RegisterController registers a controller.
func RegisterController(name string, m Controller) {
	if controller.Has(name) {
		panic("controller already exists: " + name)
	}

	logger.Info("[cms][controller] register: %s", name)
	controller.Register(name, m)
}

// GetController returns the controller by the given id.
func GetController[T any](id string) T {
	if !controller.Has(id) {
		panic("controller not registered: " + id)
	}

	s, ok := controller.MustGet(id).(T)
	if !ok {
		panic(fmt.Sprintf("controller not valid type(%v): %s", new(T), id))
	}

	return s
}

// Service returns the service container.
func (c *ControllerImpl) Service() container.Container {
	return service
}
