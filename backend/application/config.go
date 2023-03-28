package application

import (
	"fmt"

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
	Host     string   `envconfig:"host"`
	Port     string   `envconfig:"port"`
	Database DbConfig `envconfig:"primary_db"`
}

var cfg *Config

func InitializeConfig() *Config {
	if cfg != nil {
		return cfg
	}

	zap.S().Info("populating configuration from env variables")

	var temp Config
	err := envconfig.Process("backend", &temp)
	if err != nil {
		// stdlog.Fatal(fmt.Errorf("unable to process env variables: [%s]", err))
		zap.S().Fatal(fmt.Errorf("unable to process env variables: [%s]", err))
		return nil
	}

	cfg = &temp
	// zap.S().Infof("loaded configuration: %+v", cfg)

	return cfg
}
