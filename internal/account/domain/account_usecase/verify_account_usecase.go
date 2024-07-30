package account_usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/account_dto"
	"github.com/charmingruby/swrc/internal/account/domain/account_entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/logger"
)

const (
	verifyAccountUseCase = "Verify Account Use Case"
)

func (s *AccountUseCaseRegistry) VerifyAccountUseCase(dto account_dto.VerifyAccountInputDTO) error {
	solicitorHasPermission, err := s.hasRole(dto.SolicitorAccountID, account_entity.ACCOUNT_ROLE_MANAGER)
	if err != nil {
		return err
	}

	if !solicitorHasPermission {
		return core.NewUnauthorizedErr()
	}

	accToVerify, err := s.AccountRepository.FindByID(dto.AccountToVerifyID)
	if err != nil {
		return core.NewNotFoundErr("account to verify")
	}

	if err := accToVerify.Verify(dto.Verification); err != nil {
		return err
	}

	if err := s.AccountRepository.SaveVerification(accToVerify); err != nil {
		logger.LogInternalErr(verifyAccountUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
