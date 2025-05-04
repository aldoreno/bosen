package log

import (
	sglog "github.com/sourcegraph/log"
	"go.uber.org/zap"
	stdlog "log"
	"os"
)

func InitGlobalLogger() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		stdlog.Fatalf("unable to instantiate zap development logger %s", err)
	}

	zap.ReplaceGlobals(logger)
	zap.S().Info("logger set")
}

func Info() {
	log := sglog.Scoped("Log", "log component")

	// print sourcegraph log environment variables
	config := []sglog.Field{}
	for _, k := range []string{
		sglog.EnvDevelopment,
		sglog.EnvLogFormat,
		sglog.EnvLogLevel,
		sglog.EnvLogScopeLevel,
		sglog.EnvLogSamplingInitial,
		sglog.EnvLogSamplingThereafter,
	} {
		config = append(config, sglog.String(k, os.Getenv(k)))
	}
	log.Info("sglog env vars", config...)
}
