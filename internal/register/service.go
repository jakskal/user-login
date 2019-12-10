package register

import (
	"context"

	"github.com/jakskal/user-login/internal/user"
	"github.com/jakskal/user-login/pkg/rand"
	"golang.org/x/crypto/bcrypt"
)

// RepositorySystem defines operations for working with registrant data storage.
type RepositorySystem interface {
	FindByActivationCode(ctx context.Context, activationCode string) (*Registrant, error)
	Insert(ctx context.Context, registrant Registrant) error
	Update(ctx context.Context, registrant Registrant) error
}

// Service implement business operations for working with registrant.
type Service struct {
	registrantRepo RepositorySystem
	userService    user.Service
}

// NewService creates a new registrant service.
func NewService(registrantRepo RepositorySystem, userService user.Service) Service {
	return Service{
		registrantRepo: registrantRepo,
		userService:    userService,
	}
}

// Register creates an registrant.
func (s *Service) Register(ctx context.Context, req Registrant) (*string, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	activationCode := rand.String(6)
	req.Password = hashedPassword
	req.ActivationCode = activationCode
	err = s.registrantRepo.Insert(ctx, req)
	if err != nil {
		return nil, err
	}

	return &activationCode, nil
}

// Verify activate and create user.
func (s *Service) Verify(ctx context.Context, activationCode string) error {
	registrant, err := s.registrantRepo.FindByActivationCode(ctx, activationCode)
	if err != nil {
		return err
	}

	if registrant.IsActivated == true {
		return errAlreadyActivated
	}

	registrant.IsActivated = true

	user := user.User{
		Username: registrant.Username,
		Email:    registrant.Email,
		Password: registrant.Password,
	}
	_, err = s.userService.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	err = s.registrantRepo.Update(ctx, *registrant)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
