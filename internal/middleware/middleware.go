package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

// AuthWithRoleRequired scenes software that allows request to communicate and interact with application by authentication.
func AuthWithRoleRequired(permittedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		fmt.Println("token", authorizationHeader)

		if authorizationHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Token not found",
			})
			return
		} else if !strings.Contains(authorizationHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token, need bearer token",
			})
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != JWT_SIGNING_METHOD {
				return nil, fmt.Errorf("Signing method invalid")
			}

			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)

		t := time.Now()

		userID := claims["UserID"]
		role := claims["Role"]

		// Set example variable
		if userID == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid token, user id not exist",
			})
		}

		if role == nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid token, role not exist",
			})
		}

		if !isRoleHasRight(role.(string), permittedRoles...) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "forbidden request",
			})
		}

		c.Set("userID", userID)
		c.Set("role", role)
		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func isRoleHasRight(role string, roles ...string) bool {
	isHasRight := false
	for _, value := range roles {
		if role == value {
			isHasRight = true
		}
	}
	return isHasRight

}
