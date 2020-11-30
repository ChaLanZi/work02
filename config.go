package main

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	return errors.Wrap(viper.ReadInConfig(),"第一次产生错误的地方wrap")
}
