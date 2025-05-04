package log

import (
	stdlog "log"
	"os"

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

func ProvideLogger() *sglog.PostInitCallbacks {
	return sglog.Init(sglog.Resource{
		Name: manifest.AppName,
	})
}

func Logger() (sglog.Logger, func()) {
	liblog := sglog.Init(sglog.Resource{
		Name: manifest.AppName,
	})

	logger := sglog.Scoped("BosenBackend", "REST API of bosen app")
	// l := backend.Scoped("Main", "Entrypoint for the REST API")

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
	logger.Info("log configuration", config...)

	return logger, liblog.Sync
}
