package main

import (
	"github.com/hugiot/gioc/examples/ioc/provider"
	"github.com/hugiot/gioc/examples/ioc/service"
	"github.com/hugiot/gioc/src/container"
	"go.uber.org/zap"
)

func main() {
	container.AddServerProvider(&provider.AppServiceProvider{})
	container.Boot()

	logger := container.Make(service.Logger).(*zap.Logger)
	defer func() {
		_ = logger.Sync()
	}()

	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
