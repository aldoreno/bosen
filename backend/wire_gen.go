// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"bosen/application"
	"bosen/pkg/auth"
	"bosen/pkg/database"
	"bosen/pkg/user"
	"github.com/emicklei/go-restful/v3"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InjectConfig() application.Config {
	config := application.ProvideConfig()
	return config
}

func InjectDbConfig() database.DbConfig {
	config := InjectConfig()
	dbConfig := config.Database
	return dbConfig
}

func InjectContainer() *restful.Container {
	container := restful.NewContainer()
	return container
}

func InjectDiagnosticResource() *application.DiagnosticResource {
	diagnosticResource := application.NewDiagnosticResource()
	return diagnosticResource
}

func InjectLoginAction() *auth.LoginAction {
	dbConfig := InjectDbConfig()
	db := database.ProvideDatabase(dbConfig)
	userRepositoryImpl := user.NewUserRepositoryImpl(db)
	loginServiceImpl := auth.NewAuthServiceImpl(userRepositoryImpl)
	loginAction := auth.NewLoginAction(loginServiceImpl)
	return loginAction
}

func InjectAuthResource() *auth.AuthResource {
	dbConfig := InjectDbConfig()
	db := database.ProvideDatabase(dbConfig)
	userRepositoryImpl := user.NewUserRepositoryImpl(db)
	loginServiceImpl := auth.NewAuthServiceImpl(userRepositoryImpl)
	loginAction := auth.NewLoginAction(loginServiceImpl)
	authResource := auth.NewAuthResource(loginAction)
	return authResource
}

// wire.go:

func ProvideContainer() *restful.Container {
	return restful.NewContainer()
}

// UserRepositorySet provides `user.UserRepository` (an interface)
// This can be done by binding interface.
// See: https://github.com/google/wire/blob/main/docs/guide.md#binding-interfaces
var UserRepositorySet = wire.NewSet(
	InjectDbConfig, database.ProvideDatabase, user.NewUserRepositoryImpl, wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
)

var AuthServiceSet = wire.NewSet(
	UserRepositorySet, auth.NewAuthServiceImpl, wire.Bind(new(auth.LoginService), new(*auth.LoginServiceImpl)),
)

var LoginActionSet = wire.NewSet(AuthServiceSet, auth.NewLoginAction)

var AuthResourceSet = wire.NewSet(LoginActionSet, auth.NewAuthResource)