package auth

type TokenPayload struct {
	AccountID string `json:"account_id"`
	Role      string `json:"role"`
	IsValid   bool   `json:"is_valid"`
	Verified  bool   `json:"verified"`
}

type TokenService interface {
	GenerateToken(p TokenPayload) (string, error)
	ValidateToken(token string) (TokenPayload, error)
}
