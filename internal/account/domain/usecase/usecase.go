package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/port"
	"github.com/charmingruby/swrc/internal/account/domain/repository"
)

type AccountUseCase interface {
	RegisterUseCase(dto dto.RegisterInputDTO) (*dto.RegisterOutputDTO, error)
	AuthenticateUseCase(dto dto.AuthenticateInputDTO) (*dto.AuthenticateOutputDTO, error)
	VerifyAccountUseCase(dto dto.VerifyAccountInputDTO) error
	ManageAccountRoleUseCase(dto dto.ManageAccountRoleInputDTO) error
}

func NewAccountUseCaseRegistry(
	accountRepository repository.AccountRepository,
	hashPort port.HashPort,
) *AccountUseCaseRegistry {
	registry := AccountUseCaseRegistry{
		AccountRepository: accountRepository,
		HashPort:          hashPort,
	}

	return &registry
}

type AccountUseCaseRegistry struct {
	AccountRepository repository.AccountRepository
	HashPort          port.HashPort
}
