package application

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type DbConfig struct {
	Driver   string `envconfig:"driver"`
	User     string `envconfig:"user"`
	Password string `envconfig:"password"`
	Name     string `envconfig:"name"`
	Host     string `envconfig:"host"`
	Port     string `envconfig:"port"`
}

type Config struct {
	Database DbConfig `envconfig:"db"`
}

func NewConfig() *Config {
	cfg := &Config{}
	err := envconfig.Process("backend", cfg)
	if err != nil {
		zap.S().Fatal(err.Error())
	}
	zap.S().Infof("%+v", cfg)
	return cfg
}

func DatabaseConfigResolver() DbConfig {
	return NewConfig().Database
}
