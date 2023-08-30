package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	File string = "./config.toml"
)

func New() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(File)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal("init config error", err)
	}

	return v
}
