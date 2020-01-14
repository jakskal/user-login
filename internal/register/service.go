package register

import (
	"context"
	"errors"

	"github.com/jakskal/user-login/internal/customer"
	"github.com/jakskal/user-login/pkg/hash"
	"github.com/jakskal/user-login/pkg/mailer"
	"github.com/jakskal/user-login/pkg/rand"
	"github.com/jinzhu/gorm"
)

// RepositorySystem defines operations for working with registrant data storage.
type RepositorySystem interface {
	FindByActivationCode(ctx context.Context, activationCode string) (*Registrant, error)
	Insert(ctx context.Context, registrant Registrant) error
	Update(ctx context.Context, registrant Registrant) error
}

// Service implement business operations for working with registrant.
type Service struct {
	registrantRepo  RepositorySystem
	customerService customer.Service
}

// NewService creates a new registrant service.
func NewService(registrantRepo RepositorySystem, customerService customer.Service) Service {
	return Service{
		registrantRepo:  registrantRepo,
		customerService: customerService,
	}
}

// Register creates an registrant.
func (s *Service) Register(ctx context.Context, req Registrant) (*string, error) {
	hashedPassword, err := hash.Password(req.Password)
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

// Verify activate and create customer.
func (s *Service) Verify(ctx context.Context, activationCode string) error {
	registrant, err := s.registrantRepo.FindByActivationCode(ctx, activationCode)
	if err != nil {
		return err
	}

	if registrant.IsActivated == true {
		return errAlreadyActivated
	}

	registrant.IsActivated = true

	customer := customer.Customer{
		Name:     registrant.Name,
		Email:    registrant.Email,
		Password: registrant.Password,
	}
	_, err = s.customerService.CreateCustomer(ctx, customer)
	if err != nil {
		return err
	}

	err = s.registrantRepo.Update(ctx, *registrant)
	if err != nil {
		return err
	}

	return nil
}

// SendActivationEmail send email for activate account.
func (s *Service) SendActivationEmail(receiverEmail string, activationCode string) error {
	emailBody := `
	<!DOCTYPE html>
	<html>
	<body>
	<p>click this <a href="http://localhost:8080/activate/` + activationCode + `">activation link</a> to activate your account</p>
	</body>
	</html>
	`
	err := mailer.SendEmail(receiverEmail, emailBody, "Account activation")
	if err != nil {

		return err
	}

	return nil

}

// CheckRegisterPossibility check if customer can register using following email.
func (s *Service) CheckRegisterPossibility(customerEmail string) (bool, error) {
	var possible bool = true

	customer, err := s.customerService.FindCustomerByEmail(context.TODO(), customerEmail)
	if err != gorm.ErrRecordNotFound {
		return false, err
	} else if customer != nil {
		return false, errors.New("existing email already used")
	}

	return possible, nil
}
