// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jakskal/user-login/cmd/handler"
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/register"
	"github.com/jakskal/user-login/internal/token"
	"github.com/jakskal/user-login/internal/user"
	"github.com/jinzhu/gorm"
)

// Injectors from wire.go:

func initHandler(db *gorm.DB) *handler.Handler {
	repository := user.NewRepository(db)
	service := user.NewService(repository)
	controller := user.NewController(service)
	tokenService := token.NewService()
	loginService := login.NewService(repository, tokenService)
	loginController := login.NewController(loginService)
	registerRepository := register.NewRepository(db)
	registerService := register.NewService(registerRepository, service)
	registerController := register.NewController(registerService)
	handlerHandler := handler.NewHandler(controller, loginController, registerController)
	return handlerHandler
}

// wire.go:

var repositorySet = wire.NewSet(user.NewRepository, register.NewRepository, wire.Bind(new(user.RepositorySystem), new(*user.Repository)), wire.Bind(new(register.RepositorySystem), new(*register.Repository)))

var serviceSet = wire.NewSet(user.NewService, login.NewService, token.NewService, register.NewService)

var controllerSet = wire.NewSet(user.NewController, login.NewController, register.NewController)

var handlerSet = wire.NewSet(handler.NewHandler)

var appSet = wire.NewSet(
	repositorySet,
	serviceSet,
	controllerSet,
	handlerSet,
)
