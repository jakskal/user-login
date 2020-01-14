package customer

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Repository implements operations for working with customer data storage.
type Repository struct {
	db *gorm.DB
}

// NewRepository create a new customer repository
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// FindByEmail find customer by its email.
func (r *Repository) FindByEmail(ctx context.Context, customerEmail string) (*Customer, error) {

	var customer Customer
	if err := r.db.Where(&Customer{Email: customerEmail}).First(&customer).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

// FindByID find customer by its id.
func (r *Repository) FindByID(ctx context.Context, customerID string) (*Customer, error) {

	var customer Customer
	if err := r.db.First(&customer, customerID).Error; err != nil {
		return nil, err
	}

	return &customer, nil
}

// Insert a customer into data storage.
func (r *Repository) Insert(ctx context.Context, customer Customer) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Update updates a Customer.
func (r *Repository) Update(ctx context.Context, customer Customer) error {
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(customer).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// FindByCustomername find customer by its customername.
func (r *Repository) FindByCustomername(ctx context.Context, customername string) (*Customer, error) {
	var customer Customer
	err := r.db.Where("customername = ?", customername).Find(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
