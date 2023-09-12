package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// This and the rest in this package are shamelessly stolen from
// https://github.com/sourcegraph/sourcegraph/tree/30877e5b143737b008dff222861b92891d3d87f5/internal/trace

// tracerName is the name of the default tracer for the backend.
const tracerName = "bosen-backend/internal/trace"

// GetTracer returns the default tracer for the backend.
func GetTracer() oteltrace.Tracer {
	return otel.GetTracerProvider().Tracer(tracerName)
}

// Trace is a light wrapper of opentelemetry.Span. Use New to construct one.
type Trace struct {
	oteltrace.Span // never nil
}

// New returns a new Trace with the specified name in the default tracer.
// For tips on naming, see the OpenTelemetry Span documentation:
// https://opentelemetry.io/docs/specs/otel/trace/api/#span
func New(ctx context.Context, name string, attrs ...attribute.KeyValue) (Trace, context.Context) {
	return NewInTracer(ctx, GetTracer(), name, attrs...)
}

// NewInTracer is the same as New, but uses the given tracer.
func NewInTracer(ctx context.Context, tracer oteltrace.Tracer, name string, attrs ...attribute.KeyValue) (Trace, context.Context) {
	ctx, span := tracer.Start(ctx, name, oteltrace.WithAttributes(attrs...))
	return Trace{span}, ctx
}

// AddEvent records an event on this span with the given name and attributes.
//
// Note that it differs from the underlying (oteltrace.Span).AddEvent slightly, and only
// accepts attributes for simplicity.
func (t Trace) AddEvent(name string, attributes ...attribute.KeyValue) {
	t.Span.AddEvent(name, oteltrace.WithAttributes(attributes...))
}
