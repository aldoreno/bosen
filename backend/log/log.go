package log

import (
	stdlog "log"
	"os"
	"time"

	"bosen/manifest"
	sglog "github.com/sourcegraph/log"
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

func Logger() (sglog.Logger, func()) {
	liblog := sglog.Init(sglog.Resource{
		Name: manifest.AppName,
	})

	backend := sglog.Scoped("BosenBackend", "REST API of bosen app")
	l := backend.Scoped("Main", "Entrypoint for the REST API")

	// print diagnostics
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
	l.Info("log configuration", config...)

	// sample message
	l.Warn("hello world!", sglog.Time("now", time.Now()))

	return backend, liblog.Sync
}
