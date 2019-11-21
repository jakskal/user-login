package login

import (
	"context"
	"errors"

	"github.com/jakskal/user-login/internal/auth"
	"github.com/jakskal/user-login/internal/user"
	"golang.org/x/crypto/bcrypt"
)

// Service implements login service interface.
type Service struct {
	userRepo *user.Repository
	authSvc  *auth.Service
}

// NewService creates a new Login Service.
func NewService(userRepo *user.Repository, authSvc *auth.Service) Service {
	return Service{
		userRepo: userRepo,
		authSvc:  authSvc,
	}
}

func (s *Service) Login(ctx context.Context, req *LoginRequest) (*auth.Token, error) {
	user, err := s.userRepo.FindByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	v := comparePasswords(user.Password, req.Password)
	if v == false {
		return nil, errors.New("invalid password")
	}

	token, err := s.authSvc.CreateToken(ctx, &auth.CreateTokenRequest{user.ID})
	if err != nil {
		return nil, err
	}

	return token, nil

}

func comparePasswords(hashedPassword string, plainPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false
	}
	return true
}
