package login

// Request contain params to login.
type Request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
