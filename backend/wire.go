//go:build wireinject
// +build wireinject

package main

import (
	"bosen/application"
	"bosen/pkg/auth"
	"bosen/pkg/database"
	"bosen/pkg/user"

	"github.com/emicklei/go-restful/v3"
	"github.com/google/wire"
)

func InjectConfig() application.Config {
	wire.Build(application.ProvideConfig)
	return application.Config{}
}

func InjectDbConfig() database.DbConfig {
	wire.Build(
		InjectConfig,
		wire.FieldsOf(new(application.Config), "Database"),
	)
	return database.DbConfig{}
}

func ProvideContainer() *restful.Container {
	return restful.NewContainer()
}

func InjectContainer() *restful.Container {
	wire.Build(restful.NewContainer)
	return &restful.Container{}
}

func InjectDiagnosticResource() *application.DiagnosticResource {
	wire.Build(application.NewDiagnosticResource)
	return &application.DiagnosticResource{}
}

// UserRepositorySet provides `user.UserRepository` (an interface)
// This can be done by binding interface.
// See: https://github.com/google/wire/blob/main/docs/guide.md#binding-interfaces
var UserRepositorySet = wire.NewSet(
	InjectDbConfig,
	database.ProvideDatabase,
	user.NewUserRepositoryImpl,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
)

var AuthServiceSet = wire.NewSet(
	UserRepositorySet,
	auth.NewAuthServiceImpl,
	wire.Bind(new(auth.LoginService), new(*auth.LoginServiceImpl)),
)

var LoginActionSet = wire.NewSet(AuthServiceSet, auth.NewLoginAction)

var AuthResourceSet = wire.NewSet(LoginActionSet, auth.NewAuthResource)

func InjectLoginAction() *auth.LoginAction {
	wire.Build(LoginActionSet)
	return &auth.LoginAction{}
}

func InjectAuthResource() *auth.AuthResource {
	wire.Build(AuthResourceSet)
	return &auth.AuthResource{}
}
