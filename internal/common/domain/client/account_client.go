package client

import (
	"github.com/charmingruby/swrc/internal/account/domain/repository"
	"github.com/charmingruby/swrc/internal/common/core"
)

func NewAccountClient(repo repository.AccountRepository) *AccountClient {
	return &AccountClient{
		repo: repo,
	}
}

type AccountClient struct {
	repo repository.AccountRepository
}

func (c *AccountClient) AccountExists(accountID string) bool {
	if _, err := c.repo.FindByID(accountID); err != nil {
		return false
	}

	return true
}

func (c *AccountClient) ValidAccountExists(accountID string) error {
	acc, err := c.repo.FindByID(accountID)
	if err != nil {
		return core.NewNotFoundErr("account")
	}

	isValid := acc.Verification.Verified && acc.Verification.IsValid
	if !isValid {
		return core.NewUnauthorizedErr()
	}

	return nil
}

func (c *AccountClient) ValidAccountExistsAndMatchRole(accountID, role string) error {
	acc, err := c.repo.FindByID(accountID)
	if err != nil {
		return core.NewNotFoundErr("account")
	}

	isValid := acc.Verification.Verified && acc.Verification.IsValid
	if !isValid {
		return core.NewUnauthorizedErr()
	}

	if acc.Role != role {
		return core.NewUnauthorizedErr()
	}

	return nil
}
