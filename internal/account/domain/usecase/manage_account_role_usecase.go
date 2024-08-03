package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/entity"
	"github.com/charmingruby/swrc/internal/common/core"
	"github.com/charmingruby/swrc/internal/common/core/logger"
)

const (
	manageAccountRoleUseCase = "Manage Account Role Use Case"
)

func (s *AccountUseCaseRegistry) ManageAccountRoleUseCase(input dto.ManageAccountRoleInputDTO) error {
	err := s.hasPermission(input.ManagerAccountID, "manager account", entity.ACCOUNT_ROLE_MANAGER)
	if err != nil {
		return err
	}

	accToModify, err := s.AccountRepository.FindByID(input.AccountID)
	if err != nil {
		return core.NewNotFoundErr("account")
	}

	if err := accToModify.ModifyRole(input.NewRole); err != nil {
		return err
	}

	if err := s.AccountRepository.SaveRole(accToModify); err != nil {
		logger.LogInternalErr(manageAccountRoleUseCase, err)
		return core.NewInternalErr()
	}

	return nil
}
