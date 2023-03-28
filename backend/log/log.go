package log

import (
	stdlog "log"

	"go.uber.org/zap"
)

func InitGlobalLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		stdlog.Fatalf("unable to instantiate zap development logger %s", err)
	}

	zap.ReplaceGlobals(logger)
	zap.S().Info("logger set")
}
