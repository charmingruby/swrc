package factory

import (
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/account/domain/repository"
	"github.com/charmingruby/swrc/test/fake"
)

type MakeAccountInput struct {
	GithubDisplayName string
	Email             string
	Password          string
	IsValid           *bool
	Verified          *bool
	Role              string
}

func MakeAccount(
	repo repository.AccountRepository,
	in MakeAccountInput) (*entity.Account, error) {
	hashedPassword, _ := fake.NewFakeHashService().GenerateHash(in.Password)
	acc, err := entity.NewAccount(
		in.GithubDisplayName,
		in.Email,
		in.Password,
	)
	if err != nil {
		return nil, err
	}
	acc.Password = hashedPassword

	if in.Verified != nil {
		acc.Verification.Verified = *in.Verified
	}

	if in.IsValid != nil {
		acc.Verification.IsValid = *in.IsValid
	}

	if in.Role != "" {
		acc.Role = in.Role
	}

	if err := repo.Store(acc); err != nil {
		return nil, err
	}

	return acc, nil
}
