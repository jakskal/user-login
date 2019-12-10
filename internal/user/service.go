package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

// RepositorySystem defines operations for working with user data storage.
type RepositorySystem interface {
	FindByID(ctx context.Context, userID string) (*User, error)
	Insert(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	FindByUsername(ctx context.Context, username string) (*User, error)
}

// Service implement business operations for working with user.
type Service struct {
	userRepo RepositorySystem
}

// NewService creates a new user service.
func NewService(userRepo RepositorySystem) Service {
	return Service{
		userRepo: userRepo,
	}
}

// CreateUser creates an user.
func (s *Service) CreateUser(ctx context.Context, user User) (*User, error) {
	registeredUser := user
	err := s.userRepo.Insert(ctx, registeredUser)
	if err != nil {
		return nil, err
	}

	return &registeredUser, nil
}

// FindUserByID find user by its id.
func (s *Service) FindUserByID(ctx context.Context, userID string) (*User, error) {
	user, _ := s.userRepo.FindByID(ctx, userID)
	return user, nil
}

func hashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
