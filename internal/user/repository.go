package user

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Repository implements operations for working with user data storage.
type Repository struct {
	db *gorm.DB
}

// NewRepository create a new user repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FindByID find user by its id.
func (r *Repository) FindByID(ctx context.Context, userID string) (*User, error) {

	var user User
	if err := r.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Insert a user into data storage.
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

// Update updates a User.
func (r *Repository) Update(ctx context.Context, user User) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// FindByUsername find user by its username.
func (r *Repository) FindByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
