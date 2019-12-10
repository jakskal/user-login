package register

import "errors"

var (
	errAlreadyActivated      = errors.New("user is already activated")
	errInvalidActivationCode = errors.New("invalid activation code")
)

// Registrant is user that try to register.
type Registrant struct {
	ID             string `json:"-"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	ActivationCode string `json:"activation_code"`
	IsActivated    bool   `json:"is_activated"`
}
