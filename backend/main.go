package main

import (
	"bosen/application"
	"bosen/log"
	"bosen/manifest"
	"context"
	"fmt"
	"io"
	stdlog "log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.InitGlobalLogger()

	ctx := context.Background()
	shutdown, err := newTraceProvider(ctx)
	if err != nil {
		stdlog.Fatal(err)
	}
	defer func(ctx context.Context) {
		if err := shutdown(ctx); err != nil {
			stdlog.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}(ctx)

	app := application.NewApplication(
		application.WithConfig(InjectConfig()),
		application.WithContainer(InjectContainer()),
		application.WithResource(InjectDiagnosticResource()),
		application.WithResource(InjectAuthResource()),
	)

	stdlog.Fatal(app.Start(context.Background()))
}

func newStdoutExporterTracerProvider() {
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
}

func newTraceProvider(ctx context.Context) (func(context.Context) error, error) {
	res := newResource()

	// This is meant for establishing gRPC connection
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, application.GetConfig().OtlpGrpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	// Set up Trace Exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	// Register the Trace Exporter with a TracerProvider, using a batch
	// span processor to aggregate spans before export.
	tracerProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithResource(res),
		trace.WithBatcher(traceExporter),
	)
	otel.SetTracerProvider(tracerProvider)

	// Set global propagator to tracecontext (the default is no-op)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Shutdown will flush any remaining spans and shut down the exporter
	return tracerProvider.Shutdown, nil
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
