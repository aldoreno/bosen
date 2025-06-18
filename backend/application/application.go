package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	sglog "github.com/sourcegraph/log"
)

type Process interface {
	Start(context.Context)
	Stop(context.Context)
}

// Application is a container that wraps required components
type Application struct {
	log       sglog.Logger
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

	if app.log == nil {
		log.Fatal("application logger is required")
	}

	return app
}

func (a *Application) Start(ctx context.Context) error {
	a.log.Info("app starting")

	a.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.config.Host, a.config.Port),
		Handler: a.container,
	}

	// TODO: see labstack's echo implementation on starting http server
	// to be able to listen prior logging
	a.log.Info("http server started", sglog.String("server_address", a.server.Addr))

	return a.server.ListenAndServe()
}

func (a *Application) Stop(ctx context.Context) {
	a.log.Info("app stopping")
	a.server.Shutdown(ctx)
	a.log.Info("app stopped")
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
