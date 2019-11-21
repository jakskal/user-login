package user

import (
	"context"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindByID(ctx context.Context, userID string) (*User, error) {
	var user User
	r.db.First(&user)
	return &user, nil
}

func (r *Repository) Insert(ctx context.Context, user User) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *Repository) Update(ctx context.Context, user User) error {
	return nil
}

func (r *Repository) FindByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID(ctx context.Context, userID string) (*User, error)
// Insert(ctx context.Context, user User) error
// Update(ctx context.Context, user User) error
