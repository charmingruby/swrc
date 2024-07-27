package account_usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_dto"
	"github.com/charmingruby/swrc/internal/account/domain/account_repository"
)

type AccountUseCase interface {
	RegisterUseCase(dto account_dto.RegisterInputDTO) (account_dto.RegisterOutputDTO, error)
	AuthenticateUseCase(dto account_dto.AuthenticateInputDTO) (account_dto.AuthenticateOutputDTO, error)
	VerifyAccountUseCase(dto account_dto.VerifyAccountInputDTO) error
	ManageAccountRoleUseCase(dto account_dto.ManageAccountRoleInputDTO) error
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
