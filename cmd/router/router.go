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
	r := gin.Default()

	r.POST("/register", registerController.Register)
	r.POST("/login", loginController.Login)
	r.POST("/activate/:activation_code", registerController.Activate)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/users/me", userController.Me)
	}
	r.Run()
}
