// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jakskal/user-login/cmd/handler"
	"github.com/jakskal/user-login/internal/customer"
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/oauth"
	"github.com/jakskal/user-login/internal/register"
	"github.com/jakskal/user-login/internal/token"
	"github.com/jakskal/user-login/internal/user"
	"github.com/jinzhu/gorm"
)

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

// Injectors from wire.go:

func initHandler(db *gorm.DB) *handler.Handler {
	repository := user.NewRepository(db)
	service := user.NewService(repository)
	controller := user.NewController(service)
	tokenService := token.NewService()
	customerRepository := customer.NewRepository(db)
	customerService := customer.NewService(customerRepository)
	loginService := login.NewService(service, tokenService, customerService)
	loginController := login.NewController(loginService)
	registerRepository := register.NewRepository(db)
	registerService := register.NewService(registerRepository, customerService)
	registerController := register.NewController(registerService)
	oauthController := oauth.NewController(customerService)
	customerController := customer.NewController(customerService)
	handlerHandler := handler.NewHandler(controller, loginController, registerController, oauthController, customerController)
	return handlerHandler
}

// wire.go:

var repositorySet = wire.NewSet(user.NewRepository, register.NewRepository, customer.NewRepository, wire.Bind(new(user.RepositorySystem), new(*user.Repository)), wire.Bind(new(register.RepositorySystem), new(*register.Repository)), wire.Bind(new(customer.RepositorySystem), new(*customer.Repository)))

var serviceSet = wire.NewSet(user.NewService, login.NewService, token.NewService, register.NewService, customer.NewService)

var controllerSet = wire.NewSet(user.NewController, login.NewController, register.NewController, oauth.NewController, customer.NewController)

var handlerSet = wire.NewSet(handler.NewHandler)

var appSet = wire.NewSet(
	repositorySet,
	serviceSet,
	controllerSet,
	handlerSet,
)
