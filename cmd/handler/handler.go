package handler

import (
	"github.com/jakskal/user-login/internal/customer"
	"github.com/jakskal/user-login/internal/login"
	"github.com/jakskal/user-login/internal/oauth"
	"github.com/jakskal/user-login/internal/register"
	"github.com/jakskal/user-login/internal/user"
)

// Handler wrap software controller.
type Handler struct {
	UserController     *user.Controller
	LoginController    *login.Controller
	RegisterController *register.Controller
	OauthController    *oauth.Controller
	CustomerController *customer.Controller
}

// NewHandler creates a new Handler.
func NewHandler(
	userController *user.Controller,
	loginController *login.Controller,
	registerController *register.Controller,
	oauthController *oauth.Controller,
	customerController *customer.Controller,
) *Handler {
	return &Handler{
		UserController:     userController,
		LoginController:    loginController,
		RegisterController: registerController,
		OauthController:    oauthController,
		CustomerController: customerController,
	}
}
