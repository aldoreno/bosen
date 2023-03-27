package application

import (
	"context"
	"log"

	"go.uber.org/zap"
)

type Process interface {
	Start(context.Context)
	Stop(context.Context)
}

// Application is a container that wraps required components
type Application struct {
	config Config
}

func ProviderApplication(config Config) *Application {
	return &Application{config}
}

type Option func(*Application)

func NewApplicationx(opts ...Option) *Application {
	app := &Application{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// NOTE: logger should be configured the earliest
func WithLogger() Option {
	return func(_ *Application) {
		logger, err := zap.NewDevelopment()
		if err != nil {
			log.Fatalf("unable to instantiate zap development logger %s", err)
		}
		zap.ReplaceGlobals(logger)
		zap.S().Info("logger set")
	}
}

func (a *Application) AddResource(_ any) *Application {
	return a
}

func (a *Application) Start(ctx context.Context) error {
	zap.L().Info("app starting ...")

	zap.L().Info("app started")
	return nil
}

func (a *Application) Stop(ctx context.Context) {
	zap.L().Info("app stopping ...")

	zap.L().Info("app stopped")
}
