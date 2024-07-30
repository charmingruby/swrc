package account_valueobj

func NewVerification(isValid bool) *Verification {
	return &Verification{
		IsValid:  isValid,
		Verified: false,
	}
}

type Verification struct {
	IsValid  bool `json:"is_valid"`
	Verified bool `json:"verified"`
}
