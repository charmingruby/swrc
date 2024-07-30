package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/logger"
)

const (
	verifyAccountUseCase = "Verify Account Use Case"
)

func (s *AccountUseCaseRegistry) VerifyAccountUseCase(input dto.VerifyAccountInputDTO) error {
	solicitorHasPermission, err := s.hasPermission(input.SolicitorAccountID, entity.ACCOUNT_ROLE_MANAGER)
	if err != nil {
		return err
	}

	if !solicitorHasPermission {
		return core.NewUnauthorizedErr()
	}

	accToVerify, err := s.AccountRepository.FindByID(input.AccountToVerifyID)
	if err != nil {
		return core.NewNotFoundErr("account to verify")
	}

	if err := accToVerify.Verify(input.Verification); err != nil {
		return err
	}

	if err := s.AccountRepository.SaveVerification(accToVerify); err != nil {
		logger.LogInternalErr(verifyAccountUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
