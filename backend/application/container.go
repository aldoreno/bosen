package application

import (
	"github.com/emicklei/go-restful/v3"
	"go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"
)

func ProvideContainer() *restful.Container {
	container := restful.NewContainer()

	// create the Otel filter
	filter := otelrestful.OTelFilter("bosen-backend")

	// use it
	container.Filter(filter)

	return container
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
