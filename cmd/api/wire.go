//+build wireinject

package main

import (
	"github.com/jakskal/user-login/cmd/handler"
	"github.com/jakskal/user-login/internal/auth"
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/user"
	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	user.NewRepository,
	wire.Bind(new(user.UserRepository), new(*user.Repository)),
)

var serviceSet = wire.NewSet(
	user.NewService,
	login.NewService,
	auth.NewService,
)

var controllerSet = wire.NewSet(
	user.NewController,
	login.NewController,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
)
var appSet = wire.NewSet(
	repositorySet,
	serviceSet,
	controllerSet,
	handlerSet,
)

func initUserController(db *gorm.DB) *user.Controller {
	wire.Build(appSet)
	return &user.Controller{}
}

func initHandler(db *gorm.DB) *handler.Handler {
	wire.Build(appSet)
	return &handler.Handler{}
}
