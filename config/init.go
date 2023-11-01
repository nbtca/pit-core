package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfig() ServerConfig {
	v := viper.New()

	v.SetConfigFile("config.yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("config read failed", err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		fmt.Println("config unmarshal failed", err)
	}
	return serverConfig
}
