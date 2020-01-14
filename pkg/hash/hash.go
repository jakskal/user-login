package hash

import "golang.org/x/crypto/bcrypt"

// Password hashing plain password
func Password(plainPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	return string(bytes), err
}
