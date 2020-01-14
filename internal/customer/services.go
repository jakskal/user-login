package customer

import (
	"context"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// RepositorySystem defines operations for working with customer data storage.
type RepositorySystem interface {
	FindByID(ctx context.Context, customerID string) (*Customer, error)
	Insert(ctx context.Context, customer Customer) error
	Update(ctx context.Context, customer Customer) error
	FindByCustomername(ctx context.Context, customername string) (*Customer, error)
	FindByEmail(ctx context.Context, customerEmail string) (*Customer, error)
}

// Service implement business operations for working with customer.
type Service struct {
	customerRepo RepositorySystem
}

// NewService creates a new customer service.
func NewService(customerRepo RepositorySystem) Service {
	return Service{
		customerRepo: customerRepo,
	}
}

// CreateCustomer creates an customer.
func (s *Service) CreateCustomer(ctx context.Context, customer Customer) (*Customer, error) {
	registeredCustomer := customer
	err := s.customerRepo.Insert(ctx, registeredCustomer)
	if err != nil {
		return nil, err
	}

	return &registeredCustomer, nil
}

// FindCustomerByID find customer by its id.
func (s *Service) FindCustomerByID(ctx context.Context, customerID string) (*Customer, error) {
	customer, err := s.customerRepo.FindByID(ctx, customerID)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// FindCustomerByEmail find customer by its id.
func (s *Service) FindCustomerByEmail(ctx context.Context, customerEmail string) (*Customer, error) {
	customer, err := s.customerRepo.FindByEmail(ctx, customerEmail)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

// FindByEmailOrCreateCustomer find customer by email and create customer if customer not founded.
func (s *Service) FindByEmailOrCreateCustomer(ctx context.Context, req FindByEmailOrCreateCustomerRequest) (*Customer, error) {
	customer, err := s.customerRepo.FindByEmail(ctx, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if customer == nil {
		err = s.customerRepo.Insert(ctx, Customer{
			Email: req.Email,
			Name:  req.Name,
		})
		if err != nil {
			return nil, err
		}
		customer, err = s.customerRepo.FindByEmail(ctx, req.Email)
		if err != nil {
			return nil, err
		}
	}

	return customer, nil
}

func hashPassword(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
