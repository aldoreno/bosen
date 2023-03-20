package application

import (
	"log"

	"go.uber.org/zap"
)

type Application struct {
	serviceProvider map[string]interface{}
}

type Option func(*Application)

func NewApplication(opts ...Option) *Application {
	app := &Application{
		serviceProvider: make(map[string]interface{}, 0),
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func WithDatabase(resolver func() *DbConfig) Option {
	return func(a *Application) {
		a.serviceProvider["db_config"] = resolver()
		zap.S().Info("db configuration(s) set")
	}
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

func (a *Application) Start() error {
	zap.L().Info("app starting ...")
	zap.L().Info("app started")
	return nil
}

func (a *Application) Stop() {
	zap.L().Info("app stopping ...")
	zap.L().Info("app stopped")
}
