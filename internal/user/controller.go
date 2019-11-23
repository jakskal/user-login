package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller wraps an user service and implement gin context
type Controller struct {
	service Service
}

// RegisterUser regitser a user.
func (cr *Controller) RegisterUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal("failed bind sruct", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err,
		})
	}
	ctx := c.Request.Context()

	registeredUser, err := cr.service.CreateUser(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to register",
			"error":   err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "user registered succesfully",
		"username": registeredUser.Username,
	})
}

// Hello say hello
func (cr *Controller) Hello(c *gin.Context) {
	c.JSON(
		http.StatusOK, gin.H{
			"message": "you did it",
		},
	)
}

// ListUsers lists all user
func (cr *Controller) ListUsers(c *gin.Context) {
	ctx := c.Request.Context()

	user, err := cr.service.FindUserByID(ctx, "123")
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"error": "failed to get user",
			},
		)
	}
	c.JSON(
		http.StatusOK, gin.H{
			"message": "you did it",
			"user":    user,
		},
	)
}

// NewController create a new User Controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}
