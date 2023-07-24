package gormx

import (
	"fmt"

	"github.com/go-zoox/container"
	"github.com/go-zoox/logger"
)

var model = container.New()
var service = container.New()
var controller = container.New()

var creates = make(map[string]bool)

// Register registers model + service + controller.
func Register(namespace string, model Model, service Service, controller Controller) {
	if _, ok := creates[namespace]; ok {
		panic(fmt.Sprintf("[cms] app(%s) already registered", namespace))
	}
	creates[namespace] = true

	logger.Info("[cms][app] register: %s", namespace)

	if model != nil {
		RegisterModel(model.ModelName(), model)
	}

	if service != nil {
		RegisterService(service.Name(), service)
	}

	if controller != nil {
		RegisterController(controller.Name(), controller)
	}
}
