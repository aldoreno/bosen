package log

import (
	sglog "github.com/sourcegraph/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ sglog.Logger = &DummyLogger{}

type DummyLogger struct {
	*zap.Logger
}

func (z *DummyLogger) Scoped(scope string) sglog.Logger {
	return &DummyLogger{}
}

func (z *DummyLogger) With(fields ...sglog.Field) sglog.Logger {
	return &DummyLogger{}
}

func (z *DummyLogger) WithTrace(trace sglog.TraceContext) sglog.Logger {
	return &DummyLogger{}
}

func (z *DummyLogger) AddCallerSkip(skip int) sglog.Logger {
	return &DummyLogger{}
}

func (z *DummyLogger) IncreaseLevel(scope string, description string, level sglog.Level) sglog.Logger {
	return &DummyLogger{}
}

func (z *DummyLogger) WithCore(f func(c zapcore.Core) zapcore.Core) sglog.Logger {
	return &DummyLogger{}
}
