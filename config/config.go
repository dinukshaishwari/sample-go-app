package config

import (
	"sample-go-app/errors"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(name string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(name)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	errors.HandleIfError(err)
}

func GetConfig() *viper.Viper {
	return config
}
