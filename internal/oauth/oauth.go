package oauth

// CustomerInfo represent information that gained from Oauth Service
type CustomerInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	VerifiedEmail bool   `json:"verified_email"`
}
