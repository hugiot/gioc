package provider

import (
	"github.com/hugiot/gioc/examples/config"
	"github.com/hugiot/gioc/examples/ioc/service"
	"github.com/hugiot/gioc/examples/logger"
	"github.com/hugiot/gioc/src/interfaces"
	"github.com/spf13/viper"
)

type AppServiceProvider struct {
}

func (asp *AppServiceProvider) Register(c interfaces.ServiceContainer) {
	c.Single(service.Config, func(sc interfaces.ServiceContainer) any {
		return config.New()
	})
	c.Single(service.Logger, func(sc interfaces.ServiceContainer) any {
		c := sc.Make(service.Config).(*viper.Viper)
		return logger.New(c)
	})
}

func (asp *AppServiceProvider) Boot(c interfaces.ServiceContainer) {

}
