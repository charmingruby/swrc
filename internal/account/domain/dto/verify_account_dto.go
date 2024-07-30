package dto

type VerifyAccountInputDTO struct {
	SolicitorAccountID string
	AccountToVerifyID  string
	Verification       bool
}
