package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbConfig struct {
	Driver   string `envconfig:"driver"`   // Database driver: sqlite, postgresql
	User     string `envconfig:"user"`     // Username only applicable for postgresql driver
	Password string `envconfig:"password"` // Password only applicable for postgresql driver
	Name     string `envconfig:"name"`     // Database name (for sqlite it should be file name)
	Host     string `envconfig:"host"`     // Hostname only applicable for postgresql driver
	Port     string `envconfig:"port"`     // Port only applicable for postgresql driver
	Timezone string `envconfig:"timezone"` // Timezone only applicable for postgresql driver
}

func (cfg DbConfig) Postgresql() postgres.Config {
	return postgres.Config{
		// data source name, refer https://github.com/jackc/pgx
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			cfg.Host,
			cfg.User,
			cfg.Password,
			cfg.Name,
			cfg.Port,
			cfg.Timezone,
		),
	}
}

func ProvideDatabase(cfg DbConfig) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	if cfg.Driver == "" {
		zap.S().Fatalf("database driver is not provided")
		return nil
	}

	switch cfg.Driver {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.Name), &gorm.Config{})
	case "postgresql":
		db, err = gorm.Open(postgres.New(cfg.Postgresql()), &gorm.Config{})
	}

	if err != nil {
		zap.S().Fatalf("unable to open db connection: %v", err)
		return nil
	}

	return db
}
