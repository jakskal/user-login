package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jakskal/user-login/cmd/handler"
	"github.com/jakskal/user-login/internal/middleware"
)

// API generate api
func API(handler handler.Handler) {
	userController := handler.UserController
	loginController := handler.LoginController
	registerController := handler.RegisterController
	oauthController := handler.OauthController
	customerController := handler.CustomerController
	r := gin.Default()

	r.GET("/oauth/google", oauthController.HandleGoogleLoginOrRegister)
	r.GET("/callback", oauthController.HandleGoogleCallback)

	r.POST("/register", registerController.Register)
	r.GET("/activate/:activation_code", registerController.Activate)
	r.POST("/login/user", loginController.UserLogin)
	r.POST("/login/customer", loginController.CustomerLogin)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	adminAuthorized := r.Group("/admin")
	adminAuthorized.Use(middleware.AuthWithRoleRequired("ADMIN"))
	{
		adminAuthorized.POST("/users/create", userController.CreateUser)
	}

	officeAuthorized := r.Group("/")
	officeAuthorized.Use(middleware.AuthWithRoleRequired("ADMIN", "EMPLOYEE"))
	{
		officeAuthorized.GET("/users/me", userController.Me)

	}

	authorized := r.Group("/")
	authorized.Use(middleware.AuthWithRoleRequired("CUSTOMER"))
	{
		authorized.GET("/customers/me", customerController.Me)
	}
	r.Run()
}
