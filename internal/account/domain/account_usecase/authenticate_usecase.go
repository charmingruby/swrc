package account_usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_dto"
	"github.com/charmingruby/swrc/internal/common/core"
)

func (s *AccountUseCaseRegistry) AuthenticateUseCase(dto account_dto.AuthenticateInputDTO) (*account_dto.AuthenticateOutputDTO, error) {
	acc, err := s.AccountRepository.FindByEmail(dto.Email)
	if err != nil {
		return nil, core.NewInvalidCredentialsErr()
	}

	if passwordMatch := s.HashAdapter.VerifyHash(dto.Password, acc.Password); !passwordMatch {
		return nil, core.NewInvalidCredentialsErr()
	}

	return &account_dto.AuthenticateOutputDTO{
		ID:       acc.ID,
		Role:     acc.Role,
		Verified: acc.Verification.Verified,
		IsValid:  acc.Verification.IsValid,
	}, nil
}
