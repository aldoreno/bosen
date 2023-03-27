//go:build wireinject
// +build wireinject

package main

import (
	"bosen/pkg/auth"

	"github.com/google/wire"
)

func InitializeAuthSessAction() *auth.AuthenticateSessionAction {
	wire.Build(auth.NewAuthSessAction, auth.NewAuthService)
	return &auth.AuthenticateSessionAction{}
}

func InitializeAuthResource() *auth.AuthResource {
	wire.Build(InitializeAuthSessAction, auth.NewAuthResource)
	return &auth.AuthResource{}
}
