//go:build wireinject
// +build wireinject

package main

import (
	"bosen/application"
	"bosen/pkg/auth"

	"github.com/emicklei/go-restful/v3"
	"github.com/google/wire"
)

func ProvideConfig() application.Config {
	cfg := application.InitializeConfig()
	return *cfg
}

func InjectConfig() application.Config {
	wire.Build(ProvideConfig)
	return application.Config{}
}

func InjectDbConfig() (application.DbConfig, error) {
	wire.Build(
		InjectConfig,
		wire.FieldsOf(new(application.Config), "Database"),
	)
	return application.DbConfig{}, nil
}

func ProvideContainer() *restful.Container {
	return restful.NewContainer()
}

func InjectContainer() *restful.Container {
	wire.Build(restful.NewContainer)
	return &restful.Container{}
}

func InjectLoginAction() *auth.LoginAction {
	wire.Build(auth.NewLoginAction, auth.NewAuthService)
	return &auth.LoginAction{}
}

func InjectAuthResource() *auth.AuthResource {
	wire.Build(InjectLoginAction, auth.NewAuthResource)
	return &auth.AuthResource{}
}

func InjectDiagnosticResource() *application.DiagnosticResource {
	wire.Build(application.NewDiagnosticResource)
	return &application.DiagnosticResource{}
}
