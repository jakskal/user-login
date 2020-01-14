package user

// User represents user that use application.
type User struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

/*
FindByEmailOrCreateUserRequest represents paramater to find user by email
or created new user if not founded.
*/
type FindByEmailOrCreateUserRequest struct {
	Email string
	Name  string
}
