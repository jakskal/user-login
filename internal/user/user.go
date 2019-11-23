package user

// User represents user that use application.
type User struct {
	ID       string `json:"-"`
	Username string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"email"`
}
