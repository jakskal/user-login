package customer

// Customer represents customers of cupstory.
type Customer struct {
	ID       string `json:"-"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}

/*
FindByEmailOrCreateCustomerRequest represents paramater to find user by email
or created new user if not founded.
*/
type FindByEmailOrCreateCustomerRequest struct {
	Email string
	Name  string
}
