package account_usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_dto"
	"github.com/charmingruby/swrc/internal/account/domain/account_entity"
	"github.com/charmingruby/swrc/internal/common/core"
)

func (s *AccountUseCaseRegistry) RegisterUseCase(dto account_dto.RegisterInputDTO) (*account_dto.RegisterOutputDTO, error) {
	if accFoundByEmail, _ := s.AccountRepository.FindByEmail(dto.Email); accFoundByEmail != nil {
		return nil, core.NewConflictErr("account", "email")
	}

	if accFoundByGithubDisplayName, _ := s.AccountRepository.FindByGithubDisplayName(dto.GithubDisplayName); accFoundByGithubDisplayName != nil {
		return nil, core.NewConflictErr("account", "github_display_name")
	}

	passwordHash, err := s.HashAdapter.GenerateHash(dto.Password)
	if err != nil {
		return nil, err
	}

	account, err := account_entity.NewAccount(
		dto.GithubDisplayName,
		dto.Email,
		passwordHash,
	)
	if err != nil {
		return nil, err
	}

	if _, err := s.AccountRepository.Store(account); err != nil {
		return nil, core.NewInternalErr("register account use case", err.Error())
	}

	return &account_dto.RegisterOutputDTO{
		ID: account.ID,
	}, nil
}
