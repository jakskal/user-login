package login

import (
	"context"
	"errors"

	"github.com/jakskal/user-login/internal/customer"
	"github.com/jakskal/user-login/internal/token"
	"github.com/jakskal/user-login/internal/user"
	"golang.org/x/crypto/bcrypt"
)

// Service implements login service interface.
type Service struct {
	userService     user.Service
	tokenService    *token.Service
	customerService customer.Service
}

// NewService creates a new Login Service.
func NewService(userService user.Service, tokenService *token.Service, customerService customer.Service) Service {
	return Service{
		userService:     userService,
		tokenService:    tokenService,
		customerService: customerService,
	}
}

// UserLogin logging in an user.
func (s *Service) UserLogin(ctx context.Context, req *Request) (*token.Token, error) {
	user, err := s.userService.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	v := comparePasswords(user.Password, req.Password)
	if v == false {
		return nil, errors.New("invalid password")
	}

	token, err := s.tokenService.CreateToken(ctx,
		&token.CreateTokenRequest{
			UserID: user.ID,
			Role:   user.Role,
		})
	if err != nil {
		return nil, err
	}

	return token, nil

}

// CustomerLogin logging in an user.
func (s *Service) CustomerLogin(ctx context.Context, req *Request) (*token.Token, error) {
	customer, err := s.customerService.FindCustomerByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	isRegisterUsingOauth := customer.Password == ""

	if isRegisterUsingOauth {
		return nil, errors.New("customer not found")
	}

	v := comparePasswords(customer.Password, req.Password)
	if v == false {
		return nil, errors.New("invalid password")
	}

	token, err := s.tokenService.CreateToken(ctx,
		&token.CreateTokenRequest{
			UserID: customer.ID,
			Role:   "CUSTOMER",
		})
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
