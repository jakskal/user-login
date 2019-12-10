//+build wireinject

package main

import (
	"github.com/jakskal/user-login/cmd/handler"
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/register"
	"github.com/jakskal/user-login/internal/token"
	"github.com/jakskal/user-login/internal/user"
	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	user.NewRepository,
	register.NewRepository,
	wire.Bind(new(user.RepositorySystem), new(*user.Repository)),
	wire.Bind(new(register.RepositorySystem), new(*register.Repository)),
)

var serviceSet = wire.NewSet(
	user.NewService,
	login.NewService,
	token.NewService,
	register.NewService,
)

var controllerSet = wire.NewSet(
	user.NewController,
	login.NewController,
	register.NewController,
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

func initHandler(db *gorm.DB) *handler.Handler {
	wire.Build(appSet)
	return &handler.Handler{}
}
