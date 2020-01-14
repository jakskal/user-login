package token

import "github.com/dgrijalva/jwt-go"

// Token represents user tokens for authentication & authorization process (TODO).
type Token struct {
	AccessToken string `json:"access_token"`
}

// Claims represents information that contained in claim
type Claims struct {
	UserID string
	Role   string
	jwt.StandardClaims
}

type (
	// CreateTokenRequest represents parameters to create token
	CreateTokenRequest struct {
		UserID string
		Role   string
	}
)
