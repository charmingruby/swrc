package entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/account/domain/entity/account_valueobj"
	"github.com/charmingruby/swrc/internal/common/core"
)

const (
	ACCOUNT_ROLE_MANAGER   = "MANAGER"
	ACCOUNT_ROLE_DEVELOPER = "DEVELOPER"
)

func NewAccount(githubDisplayName, email, password string) (*Account, error) {
	a := Account{
		ID:                core.NewID(),
		GithubDisplayName: githubDisplayName,
		Email:             email,
		Password:          password,
		Role:              ACCOUNT_ROLE_DEVELOPER,
		Verification:      account_valueobj.NewVerification(false),
		CreatedAt:         time.Now(),
	}

	if err := core.ValidateStruct(a); err != nil {
		return nil, err
	}

	return &a, nil
}

type Account struct {
	ID                string `json:"id" validate:"required"`
	GithubDisplayName string `json:"github_display_name" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required,min=8,max=16"`
	Verification      *account_valueobj.Verification
	Role              string    `json:"role" validate:"required"`
	CreatedAt         time.Time `json:"created_at" validate:"required"`
}

func (a *Account) Verify(verification bool) error {
	if a.Verification.IsValid == verification && a.Verification.Verified {
		return core.NewValidationErr("nothing to change")
	}

	if !a.Verification.Verified {
		a.Verification.Verified = true
	}

	a.Verification.IsValid = verification

	return nil
}
