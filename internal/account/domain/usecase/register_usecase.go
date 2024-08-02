package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
)

const (
	registerUseCase = "Register Use Case"
)

func (s *AccountUseCaseRegistry) RegisterUseCase(input dto.RegisterInputDTO) (*dto.RegisterOutputDTO, error) {
	if accFoundByEmail, _ := s.AccountRepository.FindByEmail(input.Email); accFoundByEmail != nil {
		return nil, core.NewConflictErr("account", "email")
	}

	if accFoundByGithubDisplayName, _ := s.AccountRepository.FindByGithubDisplayName(input.GithubDisplayName); accFoundByGithubDisplayName != nil {
		return nil, core.NewConflictErr("account", "github_display_name")
	}

	passwordHash, err := s.HashAdapter.GenerateHash(input.Password)
	if err != nil {
		logger.LogInternalErr(registerUseCase, err)
		return nil, core.NewInternalErr()
	}

	account, err := entity.NewAccount(
		input.GithubDisplayName,
		input.Email,
		passwordHash,
	)
	if err != nil {
		return nil, err
	}

	if _, err := s.AccountRepository.Store(account); err != nil {
		logger.LogInternalErr(registerUseCase, err)
		return nil, core.NewInternalErr()
	}

	return &dto.RegisterOutputDTO{
		ID: account.ID,
	}, nil
}
