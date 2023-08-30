package services

import (
	"github.com/hugiot/gioc/examples/config"
	"github.com/hugiot/gioc/examples/logger"
	"github.com/hugiot/gioc/src/interfaces"
	"github.com/spf13/viper"
)

const (
	Config string = "config"
	Logger string = "logger"
)

type AppServiceProvider struct {
}

func (asp *AppServiceProvider) Register(c interfaces.ServiceContainer) {
	c.Single(Config, func(sc interfaces.ServiceContainer) any {
		return config.New()
	})
	c.Single(Logger, func(sc interfaces.ServiceContainer) any {
		c := sc.Make(Config).(*viper.Viper)
		return logger.New(c)
	})
}

func (asp *AppServiceProvider) Boot(c interfaces.ServiceContainer) {

}
