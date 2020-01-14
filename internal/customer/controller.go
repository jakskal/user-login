package customer

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller wraps an customer service and implement gin context
type Controller struct {
	service Service
}

// Me find loged in customer profile.
func (cr *Controller) Me(c *gin.Context) {
	ctx := c.Request.Context()
	customerID, ok := c.Get("userID")
	if !ok {
		fmt.Println("no userID")
	}
	customer, err := cr.service.FindCustomerByID(ctx, customerID.(string))
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"error": "failed to get customer",
			},
		)
		return
	}

	c.JSON(http.StatusOK, customer)
}

// CreateCustomer creates a new customer
func (cr *Controller) CreateCustomer(c *gin.Context) {
	var customer Customer
	err := c.BindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to bind struct customer",
			"error":   err.Error(),
		})
		return

	}
	ctx := c.Request.Context()

	createdCustomer, err := cr.service.CreateCustomer(ctx, customer)
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"error": "failed to get customer",
			},
		)
		return
	}

	c.JSON(http.StatusOK, createdCustomer)
}

// NewController create a new Customer Controller.
func NewController(service Service) *Controller {
	return &Controller{
		service: service,
	}
}
