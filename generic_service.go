package gormx

// Service is the interface that wraps the basic methods.
type Service interface {
	Name() string
}

// ServiceImpl is the implementation of the Service interface.
type ServiceImpl struct {
}

// Name returns the name of the service.
func (s *ServiceImpl) Name() string {
	panic("service.Name() not implemented")
}
