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
	r := gin.Default()

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		authorized.POST("/register", userController.RegisterUser)
		authorized.POST("/login", loginController.Login)
		authorized.GET("/user", userController.ListUsers)
	}
	r.Run()
}
