package handler

import (
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/register"
	"github.com/jakskal/user-login/internal/user"
)

type Handler struct {
	UserController     *user.Controller
	LoginController    *login.Controller
	RegisterController *register.Controller
}

func NewHandler(
	userController *user.Controller,
	loginController *login.Controller,
	registerController *register.Controller,
) *Handler {
	return &Handler{
		UserController:     userController,
		LoginController:    loginController,
		RegisterController: registerController,
	}
}
