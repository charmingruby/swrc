package inmemory_repository

import (
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
)

var (
	ErrNotFound = core.NewNotFoundErr("account")
)

func NewInMemoryAccountRepository() *InMemoryAccountRepository {
	return &InMemoryAccountRepository{
		Items: []entity.Account{},
	}
}

type InMemoryAccountRepository struct {
	Items []entity.Account
}

func (r *InMemoryAccountRepository) FindByID(id string) (entity.Account, error) {
	for _, acc := range r.Items {
		if acc.ID == id {
			return acc, nil
		}
	}

	return entity.Account{}, ErrNotFound
}

func (r *InMemoryAccountRepository) FindByEmail(email string) (entity.Account, error) {
	for _, acc := range r.Items {
		if acc.Email == email {
			return acc, nil
		}
	}

	return entity.Account{}, ErrNotFound
}

func (r *InMemoryAccountRepository) FindByGithubDisplayName(githubDisplayName string) (entity.Account, error) {
	for _, acc := range r.Items {
		if acc.GithubDisplayName == githubDisplayName {
			return acc, nil
		}
	}

	return entity.Account{}, ErrNotFound
}

func (r *InMemoryAccountRepository) Store(acc entity.Account) error {
	r.Items = append(r.Items, acc)
	return nil
}

func (r *InMemoryAccountRepository) SaveVerification(acc entity.Account) error {
	for idx, repoAcc := range r.Items {
		if repoAcc.ID == acc.ID {
			r.Items[idx].Verification.IsValid = acc.Verification.IsValid
			r.Items[idx].Verification.Verified = acc.Verification.Verified
			return nil
		}
	}

	return ErrNotFound
}

func (r *InMemoryAccountRepository) SaveRole(acc entity.Account) error {
	for idx, repoAcc := range r.Items {
		if repoAcc.ID == acc.ID {
			r.Items[idx].Role = acc.Role
			return nil
		}
	}

	return ErrNotFound
}
