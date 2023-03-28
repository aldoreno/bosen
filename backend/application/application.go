package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
)

type Process interface {
	Start(context.Context)
	Stop(context.Context)
}

// Application is a container that wraps required components
type Application struct {
	config    Config
	container *restful.Container
	server    *http.Server
}

type Option func(*Application)

func NewApplication(opts ...Option) *Application {
	app := &Application{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (a *Application) Start(ctx context.Context) error {
	zap.S().Info("app starting ...")

	a.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.config.Host, a.config.Port),
		Handler: a.container,
	}

	// TODO: see labstack's echo implementation on starting http server
	// to be able to listen prior logging
	zap.S().Infof("http server started on %s", a.server.Addr)

	return a.server.ListenAndServe()
}

func (a *Application) Stop(ctx context.Context) {
	zap.S().Info("app stopping ...")

	a.server.Shutdown(ctx)

	zap.S().Info("app stopped")
}
