package ioc

import (
	"fmt"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
)

type Service interface {
	Name() string
	Model() container.Container
	Service() container.Container
}

func RegisterService(name string, m Service) {
	if service.Has(name) {
		panic("service already exists: " + name)
	}

	logger.Info("[cms][service] register: %s", name)
	service.Register(name, m)
}

func GetService[T any](id string) T {
	if !service.Has(id) {
		panic("service not registered: " + id)
	}

	s, ok := service.MustGet(id).(T)
	if !ok {
		panic(fmt.Sprintf("service not valid type(%v): %s", new(T), id))
	}

	return s
}

type ServiceImpl struct {
}

func (s *ServiceImpl) Name() string {
	panic("service.Name() not implemented")
}

func (s *ServiceImpl) Model() container.Container {
	return model
}

func (s *ServiceImpl) Service() container.Container {
	return service
}
