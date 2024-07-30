package account_dto

type VerifyAccountInputDTO struct {
	SolicitorAccountID string
	AccountToVerifyID  string
	Verification       bool
}
