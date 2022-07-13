package main

import (
	"github.com/spf13/viper"
)

// Initialize init config
func configInit(prefix string) {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.SetConfigName(prefix + "_config")
	viper.AddConfigPath(".")
	_ = viper.ReadInConfig()
}
