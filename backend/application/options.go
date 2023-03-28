package application

import (
	stdlog "log"

	"github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
)

// NOTE: logger should be configured the earliest
func WithLogger() Option {
	return func(_ *Application) {
		stdlog.Println("setting up logger")

		logger, err := zap.NewDevelopment()
		if err != nil {
			stdlog.Fatalf("unable to instantiate zap development logger %s", err)
		}

		zap.ReplaceGlobals(logger)
		zap.S().Info("logger set")
	}
}

func WithConfig(cfg Config) Option {
	return func(a *Application) {
		a.config = cfg
	}
}

func WithContainer(container *restful.Container) Option {
	return func(a *Application) {
		a.container = container
	}
}

type Resource interface {
	WebService() *restful.WebService
}

func WithResource(resource Resource) Option {
	return func(a *Application) {
		a.container.Add(resource.WebService())
	}
}
