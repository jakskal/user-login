package user

import (
	"context"

	"github.com/jakskal/user-login/pkg/hash"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// RepositorySystem defines operations for working with user data storage.
type RepositorySystem interface {
	FindByID(ctx context.Context, userID string) (*User, error)
	Insert(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, userEmail string) (*User, error)
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

	hashedPassword, err := hash.Password(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = s.userRepo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindUserByID find user by its id.
func (s *Service) FindUserByID(ctx context.Context, userID string) (*User, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindUserByEmail find user by its id.
func (s *Service) FindUserByEmail(ctx context.Context, userEmail string) (*User, error) {
	user, err := s.userRepo.FindByEmail(ctx, userEmail)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmailOrCreateUser find user by email and create user if user not founded.
func (s *Service) FindByEmailOrCreateUser(ctx context.Context, req FindByEmailOrCreateUserRequest) (*User, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if user == nil {
		err = s.userRepo.Insert(ctx, User{
			Email: req.Email,
			Name:  req.Name,
		})
		if err != nil {
			return nil, err
		}
		user, err = s.userRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func hashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
