package config

import (
	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic("Config init failed:" + err.Error())
	}
}
