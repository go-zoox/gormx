package ioc

import (
	"fmt"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
)

type Controller interface {
	Name() string
	//
	Service() container.Container
	//
}

type ControllerImpl struct {
}

func RegisterController(name string, m Controller) {
	if controller.Has(name) {
		panic("controller already exists: " + name)
	}

	logger.Info("[cms][controller] register: %s", name)
	controller.Register(name, m)
}

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

func (c *ControllerImpl) Service() container.Container {
	return service
}
