package application

import (
	"github.com/emicklei/go-restful/v3"
	sglog "github.com/sourcegraph/log"
)

// NOTE: logger should be configured the earliest
func WithLogger(logger sglog.Logger) Option {
	return func(a *Application) {
		a.log = logger.Scoped("application", "main application")
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

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
