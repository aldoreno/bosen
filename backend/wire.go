//go:build wireinject
// +build wireinject

package main

import (
	"bosen/application"
	"bosen/pkg/auth"
	"bosen/pkg/auth/login"
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

func InjectContainer() *restful.Container {
	wire.Build(restful.NewContainer)
	return &restful.Container{}
}

func InjectDiagnosticResource() *application.DiagnosticResource {
	wire.Build(application.NewDiagnosticResource)
	return &application.DiagnosticResource{}
}

var DatabaseSet = wire.NewSet(
	InjectDbConfig,
	database.ProvideDatabase,
)

// UserRepositorySet provides `user.UserRepository` (an interface)
// This can be done by binding interface.
// See: https://github.com/google/wire/blob/main/docs/guide.md#binding-interfaces
var UserRepositorySet = wire.NewSet(
	DatabaseSet,
	user.NewUserRepositoryImpl,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
)

var LoginPresenterSet = wire.NewSet(
	InjectConfig,
	login.NewLoginPresenter,
	wire.Bind(new(login.LoginPresenter), new(*login.LoginPresenterImpl)),
)

var LoginServiceSet = wire.NewSet(
	UserRepositorySet,
	LoginPresenterSet,
	login.NewLoginServiceImpl,
	wire.Bind(new(login.LoginService), new(*login.LoginServiceImpl)),
)

var LoginActionSet = wire.NewSet(LoginServiceSet, login.NewLoginAction)

var AuthResourceSet = wire.NewSet(LoginActionSet, auth.NewAuthResource)

func InjectLoginAction() *login.LoginAction {
	wire.Build(LoginActionSet)
	return &login.LoginAction{}
}

func InjectAuthResource() *auth.AuthResource {
	wire.Build(AuthResourceSet)
	return &auth.AuthResource{}
}