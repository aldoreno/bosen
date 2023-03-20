package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Info(msg string, fields ...zapcore.Field) {
	zap.L().Info(msg, fields...)
}
