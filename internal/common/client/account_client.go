package client

import "github.com/charmingruby/swrc/internal/account/domain/account_usecase"

type AccountClient struct {
	service *account_usecase.AccountUseCaseRegistry
}

func NewAccountClient(svc *account_usecase.AccountUseCaseRegistry) *AccountClient {
	return &AccountClient{
		service: svc,
	}
}

func (c *AccountClient) AccountExists(accountID string) bool {
	if _, err := c.service.AccountRepository.FindByID(accountID); err != nil {
		return false
	}

	return true
}

func (c *AccountClient) IsTheAccountRole(accountID, role string) bool {
	acc, err := c.service.AccountRepository.FindByID(accountID)
	if err != nil {
		return false
	}

	if acc.Role != role {
		return false
	}

	return true
}
