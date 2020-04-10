package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type appConfig struct {
	LogLevel string `mapstructure:"log_level"`
}

// Config contain app configuration
var Config appConfig

// LoadConfig load configs from dir
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	log.Print(v.Get("log_level"))
	return v.Unmarshal(&Config)
}
