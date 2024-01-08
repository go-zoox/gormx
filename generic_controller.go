package gormx

import (
	"github.com/go-zoox/zoox"
)

// Controller is the interface that wraps the basic methods.
type Controller interface {
	Name() string
	//
	Params(ctx *zoox.Context) *Params
}

// ControllerImpl is the implementation of the Controller interface.
type ControllerImpl struct {
}

// Params returns the params.
func (c *ControllerImpl) Params(ctx *zoox.Context) *Params {
	return NewParams(ctx)
}
