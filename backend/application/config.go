package application

import (
	"bosen/pkg/database"
	"fmt"
	stdlog "log"

	"github.com/kelseyhightower/envconfig"
	sglog "github.com/sourcegraph/log"
)

type Config struct {
	Host         string            `envconfig:"host"`
	Port         string            `envconfig:"port"`
	JWTSecret    string            `envconfig:"jwt_secret"`
	Database     database.DbConfig `envconfig:"primary_db"`
	OtlpGrpcAddr string            `envconfig:"otlp_grpc_address"`
}

var cfg *Config

func GetConfig(log sglog.Logger) *Config {
	if cfg != nil {
		return cfg
	}

	log = log.Scoped("config", "application config")
	log.Info("populating configuration from env variables")

	var temp Config
	err := envconfig.Process("backend", &temp)
	if err != nil {
		stdlog.Fatal(fmt.Errorf("unable to process env variables: [%s]", err))
		return nil
	}

	cfg = &temp
	log.Debug(fmt.Sprintf("loaded configuration: %+v", cfg))

	return cfg
}

func ProvideConfig(log sglog.Logger) Config {
	cfg := GetConfig(log)
	return *cfg
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
