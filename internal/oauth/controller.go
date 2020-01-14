package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jakskal/user-login/internal/customer"
	"github.com/jakskal/user-login/internal/token"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Controller wrap oauth service
type Controller struct {
	customer customer.Service
	token    token.Service
}

var (
	oauthStateString  = os.Getenv("GOOGLE_OAUTH_STATE")
	googleOauthConfig *oauth2.Config
)

// HandleGoogleLoginOrRegister handle oauth using google acount for login or Register
func (s *Controller) HandleGoogleLoginOrRegister(c *gin.Context) {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// HandleGoogleCallback handle callback google oauth
func (s *Controller) HandleGoogleCallback(c *gin.Context) {
	customerInfo, err := getCustomerInfo(c.Request.FormValue("state"), c.Request.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get customer info from google oauth",
			"error":   err.Error(),
		})
		return
	}

	customer, err := s.FindOrCreateCustomer(context.TODO(), *customerInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to find or create customer",
			"error":   err.Error(),
		})
		return
	}

	createTokenRequest := token.CreateTokenRequest{
		UserID: customer.ID,
		Role:   "CUSTOMER",
	}
	token, err := s.token.CreateToken(context.TODO(), &createTokenRequest)

	c.JSON(http.StatusOK, token)
}

func getCustomerInfo(state string, code string) (*CustomerInfo, error) {

	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting customer info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	var customerInfo CustomerInfo
	err = json.Unmarshal(contents, &customerInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal customer info from oauth : %s", err.Error())
	}
	return &customerInfo, nil
}

// FindOrCreateCustomer find or create customer if not exist
func (s *Controller) FindOrCreateCustomer(ctx context.Context, customerInfo CustomerInfo) (*customer.Customer, error) {
	findOrCreateCustomerParams := customer.FindByEmailOrCreateCustomerRequest{
		Email: customerInfo.Email,
		Name:  customerInfo.Name,
	}
	customer, err := s.customer.FindByEmailOrCreateCustomer(ctx, findOrCreateCustomerParams)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

// NewController create a new oauth Controller.
func NewController(customer customer.Service) *Controller {
	return &Controller{
		customer: customer,
	}
}
