package main

import (
	"bosen/application"
	"bosen/log"
	"bosen/manifest"
	"context"
	"io"
	stdlog "log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	log.InitGlobalLogger()

	// Write telemetry data to a file
	f, err := os.Create("traces.txt")
	if err != nil {
		stdlog.Fatal(err)
	}
	defer f.Close()

	exporter, err := newExporter(f)
	if err != nil {
		stdlog.Fatal(err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			stdlog.Fatal(err)
		}
	}()
	otel.SetTracerProvider(traceProvider)

	app := application.NewApplication(
		application.WithConfig(InjectConfig()),
		application.WithContainer(InjectContainer()),
		application.WithResource(InjectDiagnosticResource()),
		application.WithResource(InjectAuthResource()),
	)

	stdlog.Fatal(app.Start(context.Background()))
}

func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo
		// stdouttrace.WithoutTimestamps(),
	)
}

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(manifest.AppName),
			semconv.ServiceVersion(manifest.AppVersion),
			attribute.String("environment", manifest.ReleaseVersion),
		),
	)
	return r
}
