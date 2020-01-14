package register

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Controller wraps an registrant service and implement gin context
type Controller struct {
	service Service
}

// Register regitser a registrant.
func (cr *Controller) Register(c *gin.Context) {
	var registrant Registrant
	err := c.BindJSON(&registrant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct",
			"error":   err.Error(),
		})
		return

	}

	ctx := c.Request.Context()

	canRegister, err := cr.service.CheckRegisterPossibility(registrant.Email)
	fmt.Println("erooor not found", err == gorm.ErrRecordNotFound)
	if (err != nil && err != gorm.ErrRecordNotFound) || !canRegister {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to register, impossible to register",
			"error":   err.Error(),
		})
		return
	}

	activationCode, err := cr.service.Register(ctx, registrant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to register",
			"error":   err.Error(),
		})
		return
	}

	err = cr.service.SendActivationEmail(registrant.Email, *activationCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to send activation email",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":         "register succes, please verify its you by use the following activation code",
		"activation code": activationCode,
	})
}

// Activate activate and create user.
func (cr *Controller) Activate(c *gin.Context) {
	ctx := c.Request.Context()
	activationCode := c.Param("activation_code")

	err := cr.service.Verify(ctx, activationCode)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to verify registrant",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(
		http.StatusOK, gin.H{
			"message": "verification success",
		},
	)
}

// NewController create a new User Controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}
