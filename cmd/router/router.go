package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jakskal/user-login/cmd/handler"
)

// API generate api
func API(handler handler.Handler) {
	userController := handler.UserController
	loginController := handler.LoginController
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/register", userController.RegisterUser)
	r.GET("/user", userController.ListUsers)
	r.POST("/login", loginController.Login)
	r.Run()
}
