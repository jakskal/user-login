package login

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller wraps an login service and implement gin context.
type Controller struct {
	service Service
}

// NewController create a new login controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

// UserLogin logging in an user.
func (cr *Controller) UserLogin(c *gin.Context) {
	var loginRequest *Request
	err := c.Bind(&loginRequest)
	if err != nil {
		log.Fatal("failed bind sruct", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	response, err := cr.service.UserLogin(ctx, loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// CustomerLogin logging in an user.
func (cr *Controller) CustomerLogin(c *gin.Context) {
	var loginRequest *Request
	err := c.Bind(&loginRequest)
	if err != nil {
		log.Fatal("failed bind sruct", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	response, err := cr.service.CustomerLogin(ctx, loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
