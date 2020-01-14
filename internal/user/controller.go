package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller wraps an user service and implement gin context
type Controller struct {
	service Service
}

// Me find loged in user profile.
func (cr *Controller) Me(c *gin.Context) {
	ctx := c.Request.Context()
	userID, ok := c.Get("userID")
	if !ok {
		fmt.Println("no userID")
	}
	user, err := cr.service.FindUserByID(ctx, userID.(string))
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"error": "failed to get user",
			},
		)
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser creates a new user
func (cr *Controller) CreateUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct user",
			"error":   err.Error(),
		})
		return

	}
	ctx := c.Request.Context()

	createdUser, err := cr.service.CreateUser(ctx, user)
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"error": "failed to get user",
			},
		)
		return
	}

	c.JSON(http.StatusOK, createdUser)
}

// NewController create a new User Controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}
