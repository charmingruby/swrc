package security

type Payload struct {
	AccountID string `json:"account_id"`
	Role      string `json:"role"`
	IsValid   bool   `json:"is_valid"`
	Verified  bool   `json:"verified"`
}

type Token interface {
	GenerateToken(p Payload) (string, error)
	ValidateToken(token string) bool
}
