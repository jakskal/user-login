package register

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Repository implements operations for working with registrant data storage.
type Repository struct {
	db *gorm.DB
}

// NewRepository create a new registrant repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Insert a registrant into data storage.
func (r *Repository) Insert(ctx context.Context, registrant Registrant) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&registrant).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Update updates a Registrant.
func (r *Repository) Update(ctx context.Context, registrant Registrant) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(registrant).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// FindByActivationCode find registrant by its activationCode.
func (r *Repository) FindByActivationCode(ctx context.Context, activationCode string) (*Registrant, error) {
	var registrant Registrant
	err := r.db.Where("activation_code = ?", activationCode).Find(&registrant).Error
	if err != nil {
		return nil, err
	}
	return &registrant, nil
}
