// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jakskal/user-login/cmd/handler"
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/token"
	"github.com/jakskal/user-login/internal/user"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Injectors from wire.go:

func initUserController(db *gorm.DB) *user.Controller {
	repository := user.NewRepository(db)
	service := user.NewService(repository)
	controller := user.NewController(service)
	return controller
}

func initHandler(db *gorm.DB) *handler.Handler {
	repository := user.NewRepository(db)
	service := user.NewService(repository)
	controller := user.NewController(service)
	tokenService := token.NewService()
	loginService := login.NewService(repository, tokenService)
	loginController := login.NewController(loginService)
	handlerHandler := handler.NewHandler(controller, loginController)
	return handlerHandler
}

// wire.go:

var repositorySet = wire.NewSet(user.NewRepository, wire.Bind(new(user.RepositorySystem), new(*user.Repository)))

var serviceSet = wire.NewSet(user.NewService, login.NewService, token.NewService)

var controllerSet = wire.NewSet(user.NewController, login.NewController)

var handlerSet = wire.NewSet(handler.NewHandler)

var appSet = wire.NewSet(
	repositorySet,
	serviceSet,
	controllerSet,
	handlerSet,
)
