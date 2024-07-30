package dto

type AuthenticateInputDTO struct {
	Email    string
	Password string
}

type AuthenticateOutputDTO struct {
	ID       string
	Role     string
	Verified bool
	IsValid  bool
}
