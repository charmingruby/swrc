package factory

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_entity"
	"github.com/charmingruby/swrc/internal/account/domain/account_repository"
	"github.com/charmingruby/swrc/test/fake"
)

type MakeAccountInput struct {
	GithubDisplayName string
	Email             string
	Password          string
}

func MakeAccount(
	repo account_repository.AccountRepository,
	in MakeAccountInput) (*account_entity.Account, error) {
	hashedPassword, _ := fake.NewFakeHashService().GenerateHash(in.Password)
	acc, err := account_entity.NewAccount(
		in.GithubDisplayName,
		in.Email,
		hashedPassword,
	)
	if err != nil {
		return nil, err
	}

	if _, err := repo.Store(acc); err != nil {
		return nil, err
	}

	return acc, nil
}
