package account_dto

type AuthenticateInputDTO struct {
	Email    string
	Password string
}

type AuthenticateOutputDTO struct {
	ID string
}
