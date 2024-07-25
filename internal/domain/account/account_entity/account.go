package account_entity

import (
	"time"

	"github.com/charmingruby/swrc/internal/core"
	"github.com/oklog/ulid/v2"
)

func NewAccount(githubDisplayName, email, password string) (*Account, error) {
	a := Account{
		ID:                ulid.Make().String(),
		GithubDisplayName: githubDisplayName,
		Email:             email,
		Password:          password,
		CreatedAt:         time.Now(),
	}

	if err := core.ValidateStruct(a); err != nil {
		return nil, err
	}

	return &a, nil
}

type Account struct {
	ID                string    `json:"id" validate:"required"`
	GithubDisplayName string    `json:"github_display_name" validate:"required"`
	Email             string    `json:"email" validate:"required,email"`
	Password          string    `json:"password" validate:"required,min=8,max=16"`
	CreatedAt         time.Time `json:"created_at" validate:"required"`
}
