package o11y

import (
	"bosen/manifest"
	"context"
	"go.opentelemetry.io/contrib/propagators/autoprop"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func InitTracerProvider() (*sdktrace.TracerProvider, error) {
	// Create stdout exporter to be able to retrieve
	// the collected spans.
	exporter, err := otlptrace.New(context.Background(), otlptracehttp.NewClient())
	if err != nil {
		return nil, err
	}

	// For the demonstration, use sdktrace.AlwaysSample sampler to sample all traces.
	// In a production application, use sdktrace.ProbabilitySampler with a desired probability.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(newResource()),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},

		// Configures the global TextMapPropagator to use the
		// W3C Trace Context format for propagating trace context.
		// This configuration ensures that spans have the correct
		// parent-child relationship within a trace.
		autoprop.NewTextMapPropagator(),
	))

	return tp, nil
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

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
