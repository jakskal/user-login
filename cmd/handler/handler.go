package handler

import (
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/user"
)

type Handler struct {
	UserController  *user.Controller
	LoginController *login.Controller
}

func NewHandler(
	userController *user.Controller,
	loginController *login.Controller,
) *Handler {
	return &Handler{
		UserController:  userController,
		LoginController: loginController,
	}
}
