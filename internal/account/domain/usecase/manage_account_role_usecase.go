package usecase

import (
	"github.com/charmingruby/swrc/internal/account/domain/dto"
	"github.com/charmingruby/swrc/internal/account/domain/entity"
)

func (s *AccountUseCaseRegistry) ManageAccountRoleUseCase(input dto.ManageAccountRoleInputDTO) error {
	err := s.hasPermission(input.ActorAccountID, "actor account", entity.ACCOUNT_ROLE_MANAGER)
	if err != nil {
		return err
	}

	return nil
}
