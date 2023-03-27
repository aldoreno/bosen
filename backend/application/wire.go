//go:build wireinject
// +build wireinject

package application

import (
	"bosen/log"

	"github.com/google/wire"
	"github.com/kelseyhightower/envconfig"
)

var cfg *Config

func initConfig() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	err := envconfig.Process("backend", cfg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Infof("%+v", cfg)
	return cfg, nil
}

func provideConfig() (Config, error) {
	cfg, err := initConfig()
	return *cfg, err
}

func injectConfig() (Config, error) {
	wire.Build(provideConfig)
	return Config{}, nil
}

func injectDbConfig() (DbConfig, error) {
	wire.Build(
		injectConfig,
		wire.FieldsOf(new(Config), "Database"),
	)
	return DbConfig{}, nil
}

var applicationProvider = wire.NewSet(
	injectConfig,
	wire.Struct(new(Application), "config"),
)

func NewApplication() (*Application, error) {
	wire.Build(applicationProvider)
	return &Application{}, nil
}
