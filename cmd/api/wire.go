//+build wireinject

package main

import (
	"github.com/jakskal/user-login/cmd/handler"
	"github.com/jakskal/user-login/internal/customer"
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/oauth"
	"github.com/jakskal/user-login/internal/register"
	"github.com/jakskal/user-login/internal/token"
	"github.com/jakskal/user-login/internal/user"
	"github.com/jinzhu/gorm"

	"github.com/google/wire"
)

var repositorySet = wire.NewSet(
	user.NewRepository,
	register.NewRepository,
	customer.NewRepository,
	wire.Bind(new(user.RepositorySystem), new(*user.Repository)),
	wire.Bind(new(register.RepositorySystem), new(*register.Repository)),
	wire.Bind(new(customer.RepositorySystem), new(*customer.Repository)),
)

var serviceSet = wire.NewSet(
	user.NewService,
	login.NewService,
	token.NewService,
	register.NewService,
	customer.NewService,
)

var controllerSet = wire.NewSet(
	user.NewController,
	login.NewController,
	register.NewController,
	oauth.NewController,
	customer.NewController,
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
