package main

import (
	"github.com/hugiot/gioc/examples/services"
	"github.com/hugiot/gioc/src/container"
	"go.uber.org/zap"
)

func main() {
	container.AddServerProvider(&services.AppServiceProvider{})
	container.Boot()

	logger := container.Make(services.Logger).(*zap.Logger)
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
