package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/common/core"
)

func (s *AccountUseCaseRegistry) AuthenticateUseCase(input dto.AuthenticateInputDTO) (*dto.AuthenticateOutputDTO, error) {
	acc, err := s.AccountRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, core.NewInvalidCredentialsErr()
	}

	if passwordMatch := s.HashPort.VerifyHash(input.Password, acc.Password); !passwordMatch {
		return nil, core.NewInvalidCredentialsErr()
	}

	return &dto.AuthenticateOutputDTO{
		ID:       acc.ID,
		Role:     acc.Role,
		Verified: acc.Verification.Verified,
		IsValid:  acc.Verification.IsValid,
	}, nil
}
