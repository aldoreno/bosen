package trace

import (
	"context"

	"github.com/sourcegraph/log"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// FromContext returns the Trace previously associated with ctx.
func FromContext(ctx context.Context) Trace {
	return Trace{oteltrace.SpanFromContext(ctx)}
}

// ID returns a trace ID, if any, found in the given context. If you need both trace and
// span ID, use trace.Context.
func ID(ctx context.Context) string {
	return Context(ctx).TraceID
}

// Context retrieves the full trace context, if any, from context - this includes
// both TraceID and SpanID.
func Context(ctx context.Context) log.TraceContext {
	if otelSpan := oteltrace.SpanContextFromContext(ctx); otelSpan.IsValid() {
		return log.TraceContext{
			TraceID: otelSpan.TraceID().String(),
			SpanID:  otelSpan.SpanID().String(),
		}
	}

	// no span found
	return log.TraceContext{}
}
