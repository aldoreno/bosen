package main

import (
	"bosen/application"
	"bosen/log"
	"context"
	stdlog "log"
)

func main() {
	log.InitGlobalLogger()

	app := application.NewApplication(
		application.WithConfig(InjectConfig()),
		application.WithContainer(InjectContainer()),
		application.WithResource(InjectDiagnosticResource()),
		application.WithResource(InjectAuthResource()),
	)

	stdlog.Fatal(app.Start(context.Background()))
}
