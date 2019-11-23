package token

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

// CreateToken create token
func (as *Service) CreateToken(ctx context.Context, req *CreateTokenRequest) (*Token, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		UserID: req.UserID,
	}

	jtoken := jwt.NewWithClaims(
		JWT_SIGNING_METHOD,
		claims,
	)

	ts, err := jtoken.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return nil, err
	}
	token := &Token{
		AccessToken: ts,
	}
	return token, nil
}
