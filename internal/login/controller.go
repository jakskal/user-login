package login

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (cr *Controller) Login(c *gin.Context) {
	var loginRequest *LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		log.Fatal("failed bind sruct", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err,
		})
	}

	ctx := c.Request.Context()

	response, err := cr.service.Login(ctx, loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to login",
			"error":   err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"token": response,
	})
}
