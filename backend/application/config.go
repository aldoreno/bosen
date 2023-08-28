package application

import (
	"bosen/pkg/database"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type Config struct {
	Host         string            `envconfig:"host"`
	Port         string            `envconfig:"port"`
	JWTSecret    string            `envconfig:"jwt_secret"`
	Database     database.DbConfig `envconfig:"primary_db"`
	OtlpGrpcAddr string            `envconfig:"otlp_grpc_address"`
}

var cfg *Config

func GetConfig() *Config {
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

func ProvideConfig() Config {
	cfg := GetConfig()
	return *cfg
}
