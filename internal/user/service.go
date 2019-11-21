package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	FindByID(ctx context.Context, userID string) (*User, error)
	Insert(ctx context.Context, user User) error
	Update(ctx context.Context, user User) error
	FindByUsername(ctx context.Context, username string) (*User, error)
}

type Service struct {
	userRepo UserRepository
}

func NewService(userRepo UserRepository) Service {
	return Service{
		userRepo: userRepo,
	}
}

func (s *Service) CreateUser(ctx context.Context, user User) (*User, error) {
	hashedPassword, err := hashPassword(user.Password)
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

func (s *Service) FindUserByID(ctx context.Context, userID string) (*User, error) {
	user, _ := s.userRepo.FindByID(ctx, "123")
	return user, nil
}

func (s *Service) Update(ctx context.Context, user User) (*User, error) {
	return nil, nil
}

func hashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
