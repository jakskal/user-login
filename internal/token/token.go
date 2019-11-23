package token

import "github.com/dgrijalva/jwt-go"

// Token represents user tokens for authentication & authorization process (TODO).
type Token struct {
	AccessToken string `json:"access_token"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type (
	CreateTokenRequest struct {
		UserID string
	}
)
