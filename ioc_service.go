package gormx

import (
	"fmt"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
)

// Service is the interface that wraps the basic methods.
type Service interface {
	Name() string
	Model() container.Container
	Service() container.Container
}

// RegisterService registers a service.
func RegisterService(name string, m Service) {
	if service.Has(name) {
		panic("service already exists: " + name)
	}

	logger.Info("[cms][service] register: %s", name)
	service.Register(name, m)
}

// GetService returns a service.
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

// ServiceImpl is the implementation of the Service interface.
type ServiceImpl struct {
}

// Name returns the name of the service.
func (s *ServiceImpl) Name() string {
	panic("service.Name() not implemented")
}

// Model returns the model container.
func (s *ServiceImpl) Model() container.Container {
	return model
}

// Service returns the service container.
func (s *ServiceImpl) Service() container.Container {
	return service
}
