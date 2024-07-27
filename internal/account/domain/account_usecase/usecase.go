package account_usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_dto"
	"github.com/charmingruby/swrc/internal/account/domain/account_repository"
)

type AccountUseCase interface {
	RegisterUseCase(dto account_dto.RegisterDTO) (string, error)
	AuthenticateUseCase(dto account_dto.AuthenticateDTO) (string, error)
	VerifyAccountUseCase(dto account_dto.VerifyAccountDTO) error
	ManageAccountRoleUseCase(dto account_dto.ManageAccountRoleDTO) error
}

func NewAccountUseCaseRegistry(accountRepository account_repository.AccountRepository) *AccountUseCaseRegistry {
	registry := AccountUseCaseRegistry{
		accountRepository: accountRepository,
	}

	return &registry
}

type AccountUseCaseRegistry struct {
	accountRepository account_repository.AccountRepository
}
